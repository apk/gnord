package main

import (
	"net/http"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	gnord "github.com/apk/httptools"
)

// See also https://github.com/mattn/go-cgiserver/blob/master/cgiserver.go

var addr = flag.String("addr", "127.0.0.1:4040", "http service address")
var docroot = flag.String("path", ".", "http root directory")
var iphead = flag.String("ip", ".", "header for remote IP")

func main() {
	flag.Parse()
	pth, err := filepath.Abs(*docroot)
	if (err != nil) {
		fmt.Printf("filepath.Abs(%v): %v\n",*docroot,err)
		return
	}
	http.HandleFunc("/", gnord.GnordHandleFunc(&gnord.GnordOpts{Path: pth, IpHeader: *iphead}))
	log.Fatal(http.ListenAndServe(*addr, nil))
}
