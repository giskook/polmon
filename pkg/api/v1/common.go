package v1

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Code string      `json:"code"`
	Desc string      `json:"desc"`
	Data interface{} `json:"data,omitempty"`
}

func encodeResponse(w http.ResponseWriter, code string, data interface{}, errmsg string) {
	gr := &Response{
		Code: code,
		Desc: errmsg,
		Data: data,
	}
	marshalJson(w, gr)
}

func marshalJson(w http.ResponseWriter, v interface{}) error {
	data, err := json.Marshal(v)
	if err != nil {
		return err
	}
	fmt.Fprint(w, string(data))
	log.Println(string(data))
	return nil
}

func commonReply(w http.ResponseWriter, status int, code string, data interface{}, err error) {
	w.WriteHeader(status)
	errMsg := ""
	if err != nil {
		errMsg = err.Error()
	}
	encodeResponse(w, code, data, errMsg)
}
