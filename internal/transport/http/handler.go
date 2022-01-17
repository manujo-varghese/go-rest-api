package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/manujo-varghese/go-rest-api/internal/transport/http/article"
)

// HAndler -stores pointer to article service
type Handler struct{
	Router *mux.Router
	Service *article.Service
}


// NewHandler - returns a pointer to a Handler
func NewHandler(service *article.Service)  *Handler{
	return &Handler{
		Service: service,
	}
	
}

//SetupRoutes - sets up all the routes to application
func (h *Handler) SetupRoutes()  {
	fmt.Println("Setting up Routes")
	h.Router = mux.NewRouter()

	h.Router.HandleFunc("api/article/{id}", h.GetArticle).Methods("GET")
	h.Router.HandleFunc("api/article.{tag}/{date}", h.GetArticleByTagDate).Methods("GET")
	h.Router.HandleFunc("api/article", h.PostArticle).Methods("POSt")
	h.Router.HandleFunc("/api/health", func (w http.ResponseWriter, r *http.Request)  {
		fmt.Fprintf(w, "I am Alive!")
		
	})
}

// GetArticle - retrives article by ID
func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil{
		fmt.Fprintf(w, "Unable to parse UNIT from ID")
	}
	article, err := h.Service.GetArticle(uint(i))

	if err != nil{
		fmt.Fprintf(w, "Error Retrieving Article by ID")
	}

	fmt.Fprintf(w, "%+v", article)
}

//GetArticlebyTagDate - get all articles with tag and date
func (h *Handler) GetArticleByTagDate(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	tag := vars["tag"]
	date := vars["date"]
	articles, err := h.Service.GetArticleByTagDate(tag, date)

	if err != nil{
		fmt.Fprintf(w, "Failed to retrieve all articles")
	}
	fmt.Fprintf(w, "%+v", articles)

}

// PostArticle - adds a new article
func (h *Handler) PostArticle(w http.ResponseWriter, r *http.Request)  {
	article, err := h.Service.PostArticle(article.Article{
		Title: "/",
	})
	if err != nil{
		fmt.Fprintf(w, "Failed to post new Article")
	}
	fmt.Fprintf(w, "%+v", article)
	
}