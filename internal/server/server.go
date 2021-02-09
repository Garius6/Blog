package server

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/Garius6/blog/internal/storage"
	"github.com/gorilla/mux"
)

type Server struct {
	Storage storage.Storage
}

func New(store storage.Storage) *Server {
	return &Server{
		Storage: store,
	}
}

func (s *Server) ConfigureRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", s.ArticleGetAllHandler)
	r.HandleFunc("/article/create", s.ArticleCreateHandler).Methods("POST")
	r.HandleFunc("/article/create", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/articleCreate.html")
	}).Methods("GET")
	r.HandleFunc("/article/{id}", s.ArticleGetHandler).Methods("GET")

	return r
}

func (s *Server) ArticleGetAllHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := s.Storage.Article().GetAll()
	if err != nil {
		http.Error(w, "Get all error", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("templates/articles.html")
	if err != nil {
		http.Error(w, "ParseFiles error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, articles)
}

func (s *Server) ArticleGetHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	tmpl, _ := template.ParseFiles("templates/article.html")
	a, err := s.Storage.Article().FindByID(ID)
	if err != nil {
		http.Error(w, "Database error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, a)
}

func (s *Server) ArticleCreateHandler(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	information := r.FormValue("information")

	ID, err := s.Storage.Article().Create(title, information)
	if err != nil {
		http.Error(w, "Bad values", http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/article/"+strconv.Itoa(ID), http.StatusSeeOther)
}
