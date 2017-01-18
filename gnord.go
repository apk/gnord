package main

import (
	"net/http"
	"net/http/cgi"
	"log"
	"os"
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
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		file := filepath.Join(pth, filepath.FromSlash(path))
		fmt.Printf("file=%v\n",file)
		ext := filepath.Ext(file)
		fmt.Printf("ext=%v\n",ext)
		if ext == ".cgi" {
			// Hide cgi files from plain view
			http.NotFound(w, r)
			return
		}

		f, e := os.Lstat(file)
		if f.Mode() & os.ModeSymlink != 0 {
			s, e := os.Readlink(file)
			if e == nil {
				fmt.Printf("Symlink to %v\n", s);
				http.Redirect(w, r, s, http.StatusSeeOther)
				return
			}
		}

		if os.IsNotExist(e) {
			cginame := file + ".cgi"
			fmt.Printf("cgi=%v\n", cginame)
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

	//http.Handle("/", http.FileServer(http.Dir(pth)))

	//http.HandleFunc("/ip", func(w http.ResponseWriter, r *http.Request) {
	//	h := cgi.Handler{
	//		Path: "./ip.cgi",
	//		Root: pth,
	//	}
	//	h.ServeHTTP(w, r)
	//})

	log.Fatal(http.ListenAndServe("127.0.0.1:4040", nil))
}
