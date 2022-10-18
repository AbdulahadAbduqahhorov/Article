package storage

import "github.com/AbdulahadAbduqahhorov/gin/Article/models"

type StorageI interface {
	CreateArticle(id string, article models.CreateArticleModel) error
	GetArticle(limit,offset int,search string) (articles []models.Article)
	GetArticleById(id string) (models.GetArticleByIdModel, error)
	UpdateArticle(article models.UpdateArticleModel) error
	DeleteArticle(id string) error

	CreateAuthor(id string,author models.CreateAuthorModel) error
	GetAuthor(limit,offset int,search string) (authors []models.Author)
	GetAuthorById(id string) (models.Author, error)
	UpdateAuthor(author models.UpdateAuthorModel) (error)
	DeleteAuthor(id string) (error) 
}
