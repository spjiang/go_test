// http_server.go
package main

import (
	//"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world!"))
	})


	http.ListenAndServe("127.0.0.1:8888", nil)
}