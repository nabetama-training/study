package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Cant't receive")
			break
		}

		fmt.Println("Receive back from client " + reply)

		msg := "Received : " + reply
		fmt.Println("Sending to client: " + msg)

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Cant't send")
			break
		}
	}
}

func main() {
	http.Handle("/", websocket.Handler(Echo))
	http.HandleFunc("/echo",
		func(w http.ResponseWriter, r *http.Request) {
			s := websocket.Server{Handler: websocket.Handler(Echo)}
			s.ServeHTTP(w, r)
		})
	if err := http.ListenAndServe(":9999", nil); err != nil {
		log.Fatal("Listen Serve	", err)
	}
}
