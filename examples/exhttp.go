package main

//File  : exhttp.go
//Author: Simon
//Describe: describle your function
//Date  : 2020/12/7

import (
	"fmt"
	"github.com/SimonWang00/goeureka"
	"net/http"
)

func main()  {
	goeureka.RegisterClient("http://127.0.0.1:8761","myapp", "8080", "43")
	http.HandleFunc("/hello", func(responseWriter http.ResponseWriter, request *http.Request) {
		resp := "hello goeureka!"
		_, _ = responseWriter.Write([]byte(resp))
	})
	// start server
	if err := http.ListenAndServe("127.0.0.1:8000", nil); err != nil {
		fmt.Println(err)
	}
}