package main

import (
	"io"
	"net/http"

	"golang.org/x/net/websocket"
)

func echoHandler(ws *websocket.Conn) {
	io.Copy(ws, ws)

}

func main() {
	http.Handle("/echo", websocket.Handler(echoHandler))
	http.Handle("/", http.FileServer(http.Dir("./")))
	if err := http.ListenAndServe(":9999", nil); err != nil {
		panic(err)
	}
}
