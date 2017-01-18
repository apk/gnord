package main

import (
	"net/http"
	"net/http/cgi"
	"log"
	"os"
	"fmt"
	"flag"
	"path/filepath"
)

// See also https://github.com/mattn/go-cgiserver/blob/master/cgiserver.go

var addr = flag.String("addr", "127.0.0.1:4040", "http service address")
var docroot = flag.String("path", ".", "http root directory")

func main() {
	flag.Parse()
	pth, err := filepath.Abs(*docroot)
	if (err != nil) {
		fmt.Printf("filepath.Abs(%v): %v\n",*docroot,err)
		return
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		file := filepath.Join(pth, filepath.FromSlash(path))
		ext := filepath.Ext(file)
		if ext == ".cgi" {
			// Hide cgi files from plain view
			http.NotFound(w, r)
			return
		}

		f, e := os.Lstat(file)
		if e == nil && (f.Mode() & os.ModeSymlink != 0) {
			s, e := os.Readlink(file)
			if e == nil {
				http.Redirect(w, r, s, http.StatusSeeOther)
				return
			}
		}

		if os.IsNotExist(e) {
			cginame := file + ".cgi"
			_, e = os.Stat(cginame)
			if (e == nil) {
				h := cgi.Handler{
					Path: cginame,
					Root: pth,
				}
				h.ServeHTTP(w, r)
				return
			}
		}
		http.ServeFile(w, r, file)
	})

	log.Fatal(http.ListenAndServe(*addr, nil))
}
