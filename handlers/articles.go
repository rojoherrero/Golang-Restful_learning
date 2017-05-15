package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rojoherrero/learning_go_web/services"
)

// GetAllArticles comment
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnAllArticles")
	allArticles := services.RetrieveAllArticles()
	json.NewEncoder(w).Encode(allArticles)
}

// GetArticleByID comment
func GetArticleByID(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnSingleArticle")
	vars := mux.Vars(r)
	stringKey := vars["id"]
	articleID, _ := strconv.Atoi(stringKey)
	article := services.RetrieveArticleByID(articleID)

	json.NewEncoder(w).Encode(article)

}
