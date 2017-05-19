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

// SaveArticle handler for persists a new article
func SaveArticle(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var article services.Article
	err := decoder.Decode(&article)
	log.Print(article)
	if err != nil {
		log.Print(err.Error())
	}
	id := article.SaveArticle()
	json.NewEncoder(w).Encode(HTTPResp{http.StatusOK, "Aricle created successfully", strconv.Itoa(int(id))})
}
