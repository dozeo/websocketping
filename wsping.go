package main

import (
	"crypto/tls"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("ERROR: websocket url is missing")
		os.Exit(1)
	} else {
		status := check(os.Args[1])
		if status != nil {
			fmt.Printf("ERROR: %s\n", status.Error())
			os.Exit(2)
		}
	}
}

func check(url string) error {
	var dialer websocket.Dialer
	config := tls.Config{}
	config.InsecureSkipVerify = true
	dialer.TLSClientConfig = &config
	ws, _, err := dialer.Dial(url, nil)
	if err != nil {
		return err
	}
	defer ws.Close()
	ws.SetReadDeadline(time.Now().Add(1000 * time.Millisecond))
	start := time.Now().UnixNano()
	done := make(chan error)
	ws.SetPongHandler(func(d string) error {
		end := time.Now().UnixNano()
		fmt.Printf("%d\n", (end-start)/1000/1000)
		done <- nil
		return nil
	})
	err = ws.WriteControl(websocket.PingMessage, []byte("ping"), time.Now().Add(1000*time.Millisecond))
	if err != nil {
		return err
	}
	go func() {
		var err error
		err = nil
		for err == nil {
			_, _, err = ws.ReadMessage()
		}
		done <- err
	}()
	return <-done
}
