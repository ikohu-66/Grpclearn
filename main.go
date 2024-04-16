package main

import (
	"Grpclearn/Server"
	"flag"
	"log"
)

var port string

func main() {
	err := Server.Runserver(port)
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func init() {
	flag.StringVar(&port, "p", "8000", "启动端口号")
	flag.Parse()
}
