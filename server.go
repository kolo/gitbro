package main

import (
	"html/template"
	"log"
	"net/http"
)

type Server struct {
	repo        *Repository
	contentTmpl *template.Template
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

	return s, nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	branches, err := s.repo.Branches()
	if err != nil {
		log.Println(err)
		return
	}

	err = s.contentTmpl.ExecuteTemplate(w, "base", branches)
	if err != nil {
		log.Println(err)
	}
}
