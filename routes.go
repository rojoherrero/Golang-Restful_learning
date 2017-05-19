package main

import (
	"net/http"

	"github.com/rojoherrero/learning_go_web/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"AllArticles", http.MethodGet, "/all", handlers.GetAllArticles},
	Route{"GetArticleByID", http.MethodGet, "/article/{id}", handlers.GetArticleByID},
	Route{"SaveArticle", http.MethodPost, "/article/save", handlers.SaveArticle},
}
