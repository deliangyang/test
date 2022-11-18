package main

import (
	"fmt"
	"net"
	"net/http"
	"net/http/fcgi"
	"strings"
)

type FCServer struct{}

func (F FCServer) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	m := fcgi.ProcessEnv(request)
	for k, v := range m {
		fmt.Println(k, " => ", v)
	}
	fmt.Println(strings.Repeat("-", 80))
	_, _ = writer.Write([]byte("hello world"))
}

func main() {

	listener, _ := net.Listen("tcp", "127.0.0.1:9001")

	_ = fcgi.Serve(listener, FCServer{})

}
