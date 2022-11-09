package storage

import "github.com/AbdulahadAbduqahhorov/gin/Article/models"

type StorageI interface {
	Article() ArticleRepoI
	Author() AuthorRepoI
}

type ArticleRepoI interface {
	CreateArticle(id string, article models.CreateArticleModel) error
	GetArticle(limit, offset int, search string) ([]models.Article, error)
	GetArticleById(id string) (models.GetArticleByIdModel, error)
	UpdateArticle(article models.UpdateArticleModel) error
	DeleteArticle(id string) error
}

type AuthorRepoI interface {
	CreateAuthor(id string, author models.CreateAuthorModel) error
	GetAuthor(limit, offset int, search string) ([]models.Author, error)
	GetAuthorById(id string) (models.Author, error)
	UpdateAuthor(author models.UpdateAuthorModel) error
	DeleteAuthor(id string) error
}
