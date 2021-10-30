package main

import (
	"Pixy/services"
	"os"
	"os/signal"
	"time"
)

func main(){


	serverInt := os.Args[1]
	serverAddr := os.Args[2]

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go services.ServeHttp()
	go services.StartTftp()
	go services.DhcpProcess(serverAddr, serverInt)

	go func(){
		<- c
			os.Exit(0)
		}()

	for {
	time.Sleep(time.Millisecond)
	}
}