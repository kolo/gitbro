package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	dir = flag.String("dir", currentDir(), "Path to git repository")
)

func main() {
	flag.Parse()

	s, err := NewServer(*dir)
	if err != nil {
		exitWithError(err)
	}

	log.Fatal(http.ListenAndServe(":3001", s))
}

func currentDir() string {
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
