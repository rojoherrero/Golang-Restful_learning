package services

import (
	"log"
)

// Article is a basic data model
type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

// Articles is a slice of articles
type Articles []Article

// RetrieveAllArticles retrieves all the articles from DB
func RetrieveAllArticles() []Article {
	db := connectToMariaDB()
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
	return returnedArticles

}

// RetrieveArticleByID retrieves an article by its ID
func RetrieveArticleByID(articleID int) Article {
	db := connectToMariaDB()
	defer db.Close()

	var article Article

	err := db.QueryRow("SELECT * FROM articles WHERE id = ?", articleID).Scan(&article.ID, &article.Title, &article.Desc, &article.Content)

	if err != nil {
		log.Print(err.Error())
	}
	return article
}

// SaveArticle persists article on te DB
func (a *Article) SaveArticle() int64 {
	db := connectToMariaDB()
	defer db.Close()
	stmt, _ := db.Prepare("INSERT INTO golang_test.articles (title,description,content) VALUES (?,?,?)")
	res, err := stmt.Exec(a.Title, a.Desc, a.Content)
	if err != nil {
		log.Print(err.Error())
	}
	insertedID, _ := res.LastInsertId()
	return insertedID
}
