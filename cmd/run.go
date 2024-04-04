/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/giskook/polmon/pkg/api"
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
	runCmd.PersistentFlags().Int("sync.block", 0, "sync start block")
	runCmd.PersistentFlags().Int("sync.topic", 0, "sync start block")

	runCmd.PersistentFlags().String("http.addr", "", "http server addr")
	runCmd.PersistentFlags().String("http.read_timeout", "60s", "http server read timeout")
	runCmd.PersistentFlags().String("http.write_timeout", "60s", "http server write timeout")
	runCmd.PersistentFlags().String("http.idle_timeout", "60s", "http server write timeout")
	runCmd.PersistentFlags().String("http.shudown_timeout", "60s", "http server write timeout")
}

func run() {
	httpServer := api.NewServer(api.Configure{
		Addr:            viper.GetString("http.addr"),
		WriteTimeout:    viper.GetDuration("http.write_timeout"),
		ReadTimeout:     viper.GetDuration("http.read_timeout"),
		IdleTimeout:     viper.GetDuration("http.idle_timeout"),
		ShutdownTimeout: viper.GetDuration("http.shudown_timeout"),
	})
	httpServer.Start()
}
