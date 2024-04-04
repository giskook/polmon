/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/giskook/polmon/internal/persistence/sqlite"
	"github.com/giskook/polmon/internal/statistics"
	"github.com/giskook/polmon/internal/sync"
	"github.com/giskook/polmon/pkg/api"
	v1 "github.com/giskook/polmon/pkg/api/v1"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		run()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// runCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// runCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	runCmd.PersistentFlags().String("http.addr", "0.0.0.0:8080", "http server addr")
	runCmd.PersistentFlags().String("http.read_timeout", "60s", "http server read timeout")
	runCmd.PersistentFlags().String("http.write_timeout", "60s", "http server write timeout")
	runCmd.PersistentFlags().String("http.idle_timeout", "60s", "http server write timeout")
	runCmd.PersistentFlags().String("http.shutdown_timeout", "60s", "http server write timeout")

	viper.BindPFlag("http.addr", runCmd.PersistentFlags().Lookup("http.addr"))
	viper.BindPFlag("http.read_timeout", runCmd.PersistentFlags().Lookup("http.read_timeout"))
	viper.BindPFlag("http.write_timeout", runCmd.PersistentFlags().Lookup("http.write_timeout"))
	viper.BindPFlag("http.idle_timeout", runCmd.PersistentFlags().Lookup("http.idle_timeout"))
	viper.BindPFlag("http.shutdown_timeout", runCmd.PersistentFlags().Lookup("http.shutdown_timeout"))

	runCmd.PersistentFlags().String("sqlite.path", "./fee.db", "sqlite db path")

	viper.BindPFlag("sqlite.path", runCmd.PersistentFlags().Lookup("sqlite.path"))

	runCmd.PersistentFlags().String("statistics.interval", "60s", "sqlite db path")

	viper.BindPFlag("statistics.interval", runCmd.PersistentFlags().Lookup("statistics.interval"))

	runCmd.PersistentFlags().Int("sync.block", 19548652, "sync start block")
	runCmd.PersistentFlags().String("sync.address", "0x5132A183E9F3CB7C848b0AAC5Ae0c4f0491B7aB2", "sync contract address block")
	runCmd.PersistentFlags().String("sync.topic1", "0xd1ec3a1216f08b6eff72e169ceb548b782db18a6614852618d86bb19f3f9b0d3", "sync topic 1 ")
	runCmd.PersistentFlags().String("sync.topic2", "0x0000000000000000000000000000000000000000000000000000000000000003", "sync topic 2")
	runCmd.PersistentFlags().String("sync.rpc_url", "https://rpc.ankr.com/eth", "rpc address")

	viper.BindPFlag("sync.block", runCmd.PersistentFlags().Lookup("sync.block"))
	viper.BindPFlag("sync.address", runCmd.PersistentFlags().Lookup("sync.address"))
	viper.BindPFlag("sync.topic1", runCmd.PersistentFlags().Lookup("sync.topic1"))
	viper.BindPFlag("sync.topic2", runCmd.PersistentFlags().Lookup("sync.topic2"))
	viper.BindPFlag("sync.rpc_url", runCmd.PersistentFlags().Lookup("sync.rpc_url"))
}

func run() {
	store := sqlite.NewPersistence(sqlite.Configure{Path: viper.GetString("sqlite.path")})
	handler := v1.NewHandlerV1(store)
	httpServer := api.NewServer(api.Configure{
		Addr:            viper.GetString("http.addr"),
		WriteTimeout:    viper.GetDuration("http.write_timeout"),
		ReadTimeout:     viper.GetDuration("http.read_timeout"),
		IdleTimeout:     viper.GetDuration("http.idle_timeout"),
		ShutdownTimeout: viper.GetDuration("http.shutdown_timeout"),
	}, handler)
	httpServer.Start()
	defer httpServer.Stop()

	statistics := statistics.NewStatistics(statistics.Configure{
		Internal: viper.GetDuration("statistics.interval"),
	}, store)
	go statistics.Start()
	defer statistics.Stop()

	sync := sync.NewSynchronizer(sync.Configure{
		RpcURL:  viper.GetString("sync.rpc_url"),
		Block:   viper.GetUint64("sync.block"),
		Address: common.HexToAddress(viper.GetString("sync.address")),
		Topic1:  common.HexToHash(viper.GetString("sync.topic1")),
		Topic2:  common.HexToHash(viper.GetString("sync.topic2")),
	}, store)
	sync.Start()
	defer sync.Stop()
}
