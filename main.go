package main

import (
	"net/http"
	"os"
)

func main() {
	for _, arg := range os.Args {
		println(arg)
	}
	http.HandleFunc("/protected", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm protected"))
	})


	http.HandleFunc("/safe", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("test"))
	})
}