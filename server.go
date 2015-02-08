package main

import (
	"log"
	"net/http"
)

type Server struct {
	repo   *Repository
	static http.Handler
}

func NewServer(path string) (*Server, error) {
	var err error
	s := &Server{}

	s.repo, err = OpenRepository(path)
	if err != nil {
		return nil, err
	}

	s.static = staticHandler{
		http.Dir("bower_components"),
		http.Dir("webroot"),
		http.Dir(".tmp"),
	}

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s: %s\n", r.Method, r.URL.Path)
	s.static.ServeHTTP(w, r)
}
