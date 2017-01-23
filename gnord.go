package main

import (
	"net/http"
	"flag"
	"fmt"
	"log"
	"path/filepath"
	gnord "github.com/apk/httptools"
)

// https://gist.github.com/denji/12b3a568f092ab951456
// See also https://github.com/mattn/go-cgiserver/blob/master/cgiserver.go

var addr = flag.String("addr", "127.0.0.1:4040", "http service address")
var docroot = flag.String("path", ".", "http root directory")
var iphead = flag.String("ip", "", "header for remote IP")
var wellknown = flag.String("well-known", "", "host for .well-known")

func main() {
	mux := http.NewServeMux()
	flag.Parse()
	pth, err := filepath.Abs(*docroot)
	if (err != nil) {
		fmt.Printf("filepath.Abs(%v): %v\n",*docroot,err)
		return
	}

	mux.HandleFunc("/", gnord.GnordHandleFunc(&gnord.GnordOpts{Path: pth, IpHeader: *iphead}))

	if *wellknown != "" {
		mux.HandleFunc("/.well-known/", gnord.SSLForwarderHandleFunc(*wellknown))
	}

	log.Fatal(http.ListenAndServe(*addr, mux))
}
