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

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := s.ConfigureRoutes()
	router.ServeHTTP(w, r)
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

	tmpl, err := template.ParseFiles("templates/article.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	a, err := s.Storage.Article().FindByID(ID)
	if err != nil {
		if err == storage.ErrNotFound {
			http.Error(w, "Not found", http.StatusNotFound)
			return
		}
		http.Error(w, "Database error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, a)
}

func (s *Server) ArticleCreateHandler(w http.ResponseWriter, r *http.Request) {
	information, title := r.FormValue("title"), r.FormValue("information")
	if information == "" || title == "" {
		http.Error(w, "Bad values", http.StatusBadRequest)
		return
	}

	ID, err := s.Storage.Article().Create(title, information)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/article/"+strconv.Itoa(ID), http.StatusSeeOther)
}
