package storage

import (
	"fmt"
	"time"

	"github.com/AbdulahadAbduqahhorov/gin/Article/models"
	"github.com/google/uuid"
)

var InMemoryArticle []models.Article

func CreateArticle(article models.CreateArticleModel) models.Article {
	var response models.Article
	t := time.Now()
	response.CreatedAt = &t
	id := uuid.New()
	response.Id = id.String()

	response.Content = article.Content
	response.AuthorId = article.AuthorId
	InMemoryArticle = append(InMemoryArticle, response)
	return response
}



func GetArticle() (articles []models.Article) {
	for _, article := range InMemoryArticle {
		if article.DeletedAt == nil {
			articles = append(articles, article)
		}
	}
	return
}

func GetArticleById(id string) (models.Article, error) {
	var article models.Article

	for _, v := range InMemoryArticle {
		if v.Id == id && v.DeletedAt == nil {
			return v, nil
		}
	}
	return article, fmt.Errorf("article not found with id %s", id)
}

func UpdateArticle(article models.UpdateArticleModel) (models.Article, error) {

	var response models.Article
	for i := range InMemoryArticle {
		if InMemoryArticle[i].Id == article.Id && InMemoryArticle[i].DeletedAt == nil {
			t := time.Now()
			response.UpdatedAt = &t
			response.CreatedAt = InMemoryArticle[i].CreatedAt
			response.Content = article.Content
			response.AuthorId = article.AuthorId
			response.Id = article.Id
			InMemoryArticle[i] = response

			return response, nil
		}
	}
	return response, fmt.Errorf("article not found with id %s", article.Id)
}

func DeleteArticle(id string) (models.Article, error) {
	article, err := GetArticleById(id)
	if err != nil {
		return article, fmt.Errorf("article not found with id %s", id)
	}
	t := time.Now()
	article.DeletedAt = &t
	return article, nil

}
