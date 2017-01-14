package main

import (
	"net/http"
	"net/http/cgi"
	"log"
	"fmt"
	"path/filepath"
)

// See also https://github.com/mattn/go-cgiserver/blob/master/cgiserver.go

func main() {
	pth, err := filepath.Abs(".")
	if (err != nil) {
		fmt.Printf("err=%v\n",err)
		return
	}
	http.Handle("/", http.FileServer(http.Dir(pth)))

	http.HandleFunc("/ip", func(w http.ResponseWriter, r *http.Request) {
		h := cgi.Handler{
			Path: "./ip.cgi",
			Root: pth,
		}
		h.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe("127.0.0.1:4040", nil))
}
