package services

import (
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

type logStruct struct{
	remoteAddress string
	uri string
}

func GetOutbound() (string ,string) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	localAddr_raw := conn.LocalAddr().(*net.UDPAddr)
	localAddr := strings.Split(localAddr_raw.String(),":")[0]
	println(localAddr)
	interfaces, _ := net.Interfaces()
	for _, netInterface := range interfaces{
		addrs, _ := netInterface.Addrs()
		for _, addr_raw := range addrs{
			addr := strings.Split(addr_raw.String(), "/")[0]
			if addr==localAddr{
				HandleLog("Default Interface: " + netInterface.Name + ":" +  localAddr)
				return netInterface.Name, localAddr
			}
		}

	}
	HandleLog("Cannot find default interface.")
	os.Exit(1)
	return "", ""
}

func serveWithLogging() http.Handler {
	logFn := func(rw http.ResponseWriter, r *http.Request) {
		logLine := buildLog(r)
		HandleLog(logLine)
		buf, err := ioutil.ReadFile("images" + r.RequestURI)
		if err != nil {
			HandleLog("IO Error: " + err.Error())
		}
		_, err = rw.Write(buf)
		if err != nil {
			HandleLog(err.Error())
		}
	}
	return http.HandlerFunc(logFn)
}

func buildLog(r *http.Request) string{
	logEntry := logStruct{
		remoteAddress: r.RemoteAddr,
		uri:           r.RequestURI,
	}
	return string(logEntry.remoteAddress + " : " + logEntry.uri)
}

func ServeHttp(){
	http.Handle("/", serveWithLogging())
	http.ListenAndServe(":80", nil)
}
