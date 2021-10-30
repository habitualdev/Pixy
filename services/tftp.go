package services

import (
	"Pixy/embed"
	"bytes"
	"fmt"
	"github.com/pin/tftp"
	"io"
	"io/ioutil"
	"strconv"
	"time"
)

// TFTP is served from the boot directory, along with HTTP

// readHandler is called when client starts file download from server
func readHandler(filename string, rf io.ReaderFrom) error {
	var file []byte
	var err error
	if filename == "pxelinux.cfg/default" {
	file, err = ioutil.ReadFile("default")
	} else {file, err = embed.F.ReadFile(filename)}
	if err != nil {
		HandleLog(err.Error())
		return err
	}
	reader := bytes.NewReader(file)
	n, err := rf.ReadFrom(reader)
	if err != nil {
		HandleLog(err.Error())
		return err
	}
	err = fmt.Errorf( "TFTP: " + strconv.Itoa(int(n)) + " Bytes Sent")
	HandleLog(err.Error())
	return err
}

func StartTftp() {
	// use nil in place of handler to disable write operations
	s := tftp.NewServer(readHandler, nil)
	s.SetTimeout(5 * time.Second) // optional
	err := s.ListenAndServe(":69") // blocks until s.Shutdown() is called
	if err != nil {
		HandleLog(err.Error())
	}
}
