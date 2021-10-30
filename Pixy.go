package main

import (
	"Pixy/services"
	"os"
	"os/signal"
	"time"
)

func main(){

	var serverInt string
	var serverAddr string

	if len(os.Args) > 1 {
		serverInt = os.Args[1]
		serverAddr = os.Args[2]
	}else {
		serverInt, serverAddr = services.GetOutbound()
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go services.ServeHttp()
	go services.StartTftp()
	go services.DhcpProcess(serverAddr, serverInt)
	println("Press Ctrl-C to exit...")
	println(services.GetOutbound())
	go func(){
		<- c
			os.Exit(0)
		}()

	for {
	time.Sleep(time.Millisecond)
	}
}