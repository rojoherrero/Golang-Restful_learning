package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the Home Page")
	log.Println("Endpotin Hit: homePage")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/all", returnAllArticles)
	myRouter.HandleFunc("/article/{id}", returnSingleArticle)
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func main() {

	handleRequests()
}

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnAllArticles")
	db := mariaDb()
	defer db.Close()
	results, err := db.Query("SELECT * FROM articles")
	if err != nil {
		log.Print(err.Error())
	}

	var returnedArticles Articles

	for results.Next() {
		var article Article
		err := results.Scan(&article.ID, &article.Title, &article.Desc, &article.Content)
		if err != nil {
			log.Print(err.Error())
		}
		returnedArticles = append(returnedArticles, article)
	}

	json.NewEncoder(w).Encode(returnedArticles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnSingleArticle")
	db := mariaDb()
	defer db.Close()
	vars := mux.Vars(r)
	stringKey := vars["id"]
	key, _ := strconv.Atoi(stringKey)

	var article Article

	err := db.QueryRow("SELECT * FROM articles WHERE id = ?", key).Scan(&article.ID, &article.Title, &article.Desc, &article.Content)

	if err != nil {
		log.Print(err.Error())
	}

	json.NewEncoder(w).Encode(article)

}

func mariaDb() *sql.DB {
	db, err := sql.Open("mysql", "root:root@/golang_test")
	if err != nil {
		panic(err.Error)
	}
	log.Println("Connection to MariaDB created")
	return db
}
