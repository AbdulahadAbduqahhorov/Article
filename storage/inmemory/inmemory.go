package inmemory

import "github.com/AbdulahadAbduqahhorov/gin/Article/models"

type InMemory struct {
	Db *DB
}

type DB struct {
	InMemoryArticle []models.Article
	InMemoryAuthor  []models.Author
}
