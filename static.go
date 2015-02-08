package main

import (
	"net/http"
	"path"
	"strings"
)

type staticHandler []http.FileSystem

func (s staticHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" && r.Method != "HEAD" {
		return
	}

	file := r.URL.Path
	for _, dir := range s {
		f, err := dir.Open(file)
		if err != nil {
			continue
		}
		defer f.Close()

		fi, err := f.Stat()
		if err != nil {
			continue
		}

		if fi.IsDir() {
			if !strings.HasSuffix(r.URL.Path, "/") {
				http.Redirect(w, r, r.URL.Path+"/", http.StatusFound)
			}

			file = path.Join(file, "index.html")
			f, err = dir.Open(file)
			if err != nil {
				continue
			}
			defer f.Close()

			fi, err = f.Stat()
			if err != nil {
				continue
			}

		}

		http.ServeContent(w, r, r.URL.Path, fi.ModTime(), f)
		return
	}

	http.Error(w, "404 File not found", http.StatusNotFound)
}
