package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	static http.Handler
}

func NewServer() *Server {
	s := &Server{}

	s.static = staticHandler{
		http.Dir("bower_components"),
		http.Dir("webapp"),
		http.Dir(".tmp"),
	}

	router := mux.NewRouter()
	router.NotFoundHandler = s.static

	router.HandleFunc("/branches", branchesHandler)

	s.router = router

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: %s\n", r.Method, r.URL.Path)
	s.router.ServeHTTP(w, r)
}

func branchesHandler(w http.ResponseWriter, r *http.Request) {
	branches, err := repo.Branches()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	value := struct {
		Branches []string `json:"branches"`
	}{branches}

	output := bytes.NewBuffer([]byte{})
	enc := json.NewEncoder(output)
	if err := enc.Encode(value); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(output.Bytes()); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
