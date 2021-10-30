package services

import (
	"io"
	"log"
	"os"
)

func HandleLog(log_line string){
	logFile, _ := os.OpenFile("nanny.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	wrt := io.MultiWriter(logFile)
	WriteToLog(wrt, log_line)

}

func WriteToLog(wrt io.Writer,log_line string){
	log.SetOutput(wrt)
	log.Println(log_line)
}
