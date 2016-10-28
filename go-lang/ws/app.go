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
	http.HandleFunc("/echo",
		func(w http.ResponseWriter, r *http.Request) {
			s := websocket.Server{Handler: websocket.Handler(echoHandler)}
			s.ServeHTTP(w, r)
		})

	http.Handle("/", http.FileServer(http.Dir("./")))
	if err := http.ListenAndServe(":9999", nil); err != nil {
		panic(err)
	}
}
