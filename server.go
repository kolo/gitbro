package main

import (
	"html/template"
	"log"
	"net/http"
)

type Server struct {
	repo        *Repository
	contentTmpl *template.Template

	static http.Handler
}

func NewServer(path string) (*Server, error) {
	var err error
	s := &Server{}

	s.repo, err = OpenRepository(path)
	if err != nil {
		return nil, err
	}

	s.contentTmpl, err = template.ParseFiles("views/base.html", "views/content.html")
	if err != nil {
		return nil, err
	}

	s.static = http.FileServer(http.Dir(""))

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var t *template.Template

	switch r.URL.Path {
	case "/":
		t = s.contentTmpl
	default:
		s.static.ServeHTTP(w, r)
		return
	}

	commits, err := s.repo.Log("refs/heads/master")
	if err != nil {
		log.Println(err)
		return
	}

	err = t.ExecuteTemplate(w, "base", commits)
	if err != nil {
		log.Println(err)
	}
}
