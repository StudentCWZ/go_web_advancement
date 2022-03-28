/*
   @Author: StudentCWZ
   @Description:
   @File: main
   @Software: GoLand
   @Project: docker_demo
   @Date: 2022/3/28 11:23
*/

package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("hello liwenzhou.com!"))
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", hello)
	server := &http.Server{
		Addr: ":8080",
	}
	fmt.Println("server startup ...")
	if err := server.ListenAndServe(); err != nil {
		fmt.Printf("server startup failed, err: %v\n", err)
		return
	}
}
