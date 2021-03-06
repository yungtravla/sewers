package main

import (
  "io/ioutil"
  "net/http"
  "time"
)

const (
  obf_const_InterpreterGetTag    = "obf_tag_interpreter_get_tag"
  obf_const_InterpreterPostTag   = "obf_tag_interpreter_post_tag"
  obf_const_InterpreterStreamTag = "obf_tag_interpreter_stream_tag"
  obf_const_Port                 = "8008"
  obf_const_RequestDir           = "obf_tag_request_dir"
  obf_const_ResponseDir          = "obf_tag_response_dir"
  obf_const_SewersGetTag         = "obf_tag_sewers_get_tag"
  obf_const_SewersPostTag        = "obf_tag_sewers_get_tag"
  obf_const_SewersStreamTag      = "obf_tag_sewers_stream_tag"
  obf_const_StreamDir            = "obf_tag_stream_dir"
)

func obf_func_serve(res http.ResponseWriter, req *http.Request) {
  defer req.Body.Close()
  obf_var_body, _ := ioutil.ReadAll(req.Body)
  println(string(obf_var_body))
}

func main() {
  obf_var_server := &http.Server {
    Addr:              ":" + obf_const_Port,
    Handler:           http.HandlerFunc(obf_func_serve),
    ReadTimeout:       10 * time.Second,
    ReadHeaderTimeout: 10 * time.Second,
    WriteTimeout:      5 * time.Second,
  }
  println(obf_var_server.ListenAndServe().Error())
}
