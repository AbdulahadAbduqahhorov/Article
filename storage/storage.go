package storage

import "github.com/AbdulahadAbduqahhorov/gin/Article/models"

type StorageI interface {
	CreateArticle(id string, article models.CreateArticleModel) error
	GetArticle() (articles []models.Article)
	GetArticleById(id string) (models.GetArticleByIdModel, error)
	UpdateArticle(article models.UpdateArticleModel) error
	DeleteArticle(id string) error

	CreateAuthor(id string,author models.CreateAuthorModel) error
	GetAuthor() (authors []models.Author)
	GetAuthorById(id string) (models.Author, error)
	UpdateAuthor(author models.UpdateAuthorModel) (models.Author, error)
	DeleteAuthor(id string) (models.Author, error) 
}
