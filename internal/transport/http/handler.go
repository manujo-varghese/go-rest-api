package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/manujo-varghese/go-rest-api/internal/article"
)

// HAndler -stores pointer to article service
type Handler struct{
	Router *mux.Router
	Service *article.Service
}

// Response - an object to store respose from api

type Response struct{
	Message string
	Error string

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

	h.Router.HandleFunc("/api/article/{id}", h.GetArticle).Methods("GET")
	h.Router.HandleFunc("/api/tags/{tag}/{date}", h.GetArticleByTagDate).Methods("GET")
	h.Router.HandleFunc("/api/article", h.PostArticle).Methods("POST")
	h.Router.HandleFunc("/api/health", func (w http.ResponseWriter, r *http.Request)  {
		if err := sendOKResponse(w,Response{Message: "I am Alive"}); err != nil{
			panic(err)
		}
		
	})
}

// GetArticle - retrives article by ID
func (h *Handler) GetArticle(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil{
		sendErrorResponse(w, "Unable to parse UNIT from ID", err)
		return
	}
	article, err := h.Service.GetArticle(uint(i))

	if err != nil{
		sendErrorResponse(w, "Error Retrieving Article by ID", err)
		return
	}
	if err := sendOKResponse(w, article); err != nil{
		panic(err)
	}
}

//GetArticlebyTagDate - get all articles with tag and date
func (h *Handler) GetArticleByTagDate(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	tag := vars["tag"]
	date := vars["date"]
	//articles, err := h.Service.GetArticleByTagDate(tag, date)
	result, err := h.Service.GetArticleByTagDate(tag, date)
	if err != nil{
		sendErrorResponse(w, "Failed to retrieve all articles", err)
		return
	}
	if err := sendOKResponse(w,result); err != nil{
		panic(err)
	}

}


// PostArticle - adds a new article
func (h *Handler) PostArticle(w http.ResponseWriter, r *http.Request)  {
	var article article.Article
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil{
		sendErrorResponse(w, "Failed to decode json Body", err)
		return
	}
	const layout = "2006-01-02"
	tm, _ := time.Parse(layout,time.Now().Format("2006-01-02"))
	article.Date = tm
	article, err := h.Service.PostArticle(article)
	if err != nil{
		sendErrorResponse(w, "Failed to post new Article", err)
		return
	}

	if err := sendOKResponse(w,article); err != nil{
		panic(err)
	}
	
}

func sendOKResponse(w http.ResponseWriter, resp interface{}) error  {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(resp)
	
}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}