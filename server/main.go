package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	dir  = flag.String("dir", currentDir(), "Path to git repository")
	repo *Repository
)

func init() {
	webroot := os.Getenv("webroot")
	if webroot != "" {
		if err := os.Chdir(webroot); err != nil {
			exitWithError(err)
		}
	}
}

func main() {
	flag.Parse()

	var err error

	repo, err = OpenRepository(*dir)
	if err != nil {
		exitWithError(err)
	}

	s := NewServer()
	log.Fatal(http.ListenAndServe(":3001", s))
}

func currentDir() string {
	dir := os.Getenv("dir")
	if dir != "" {
		return dir
	}

	dir, err := os.Getwd()
	if err != nil {
		exitWithError(err)
	}

	return dir
}

func exitWithError(err error) {
	fmt.Fprintf(os.Stderr, "error: %v\n", err)
	os.Exit(1)
}
