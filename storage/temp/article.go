package storage

import (
	"fmt"
	"time"

	"github.com/AbdulahadAbduqahhorov/gin/Article/models"
)

var InMemoryArticle []models.Article

func CreateArticle(id string, article models.CreateArticleModel) error {
	var response models.Article
	_, err := GetAuthorById(article.AuthorId)
	if err != nil {
		return err
	}
	t := time.Now()
	response.CreatedAt = &t
	response.Id = id
	response.Content = article.Content
	response.AuthorId = article.AuthorId
	InMemoryArticle = append(InMemoryArticle, response)

	return nil

}

func GetArticle() (articles []models.Article) {
	for _, article := range InMemoryArticle {
		if article.DeletedAt == nil {
			articles = append(articles, article)
		}
	}
	return
}

func GetArticleById(id string) (models.GetArticleByIdModel, error) {
	var article models.GetArticleByIdModel

	for _, v := range InMemoryArticle {
		if v.Id == id && v.DeletedAt == nil {
			author, err := GetAuthorById(v.AuthorId)
			if err != nil {
				return article, err
			}
			article.Id = v.Id
			article.Content = v.Content
			article.Author = author
			article.CreatedAt = v.CreatedAt
			article.UpdatedAt = v.UpdatedAt
			article.DeletedAt = v.DeletedAt
			return article, nil
		}
	}
	return article, fmt.Errorf("article not found with id %s", id)
}

func UpdateArticle(article models.UpdateArticleModel) (error) {

	
	for i,v := range InMemoryArticle {
		if InMemoryArticle[i].Id == article.Id && InMemoryArticle[i].DeletedAt == nil {
			t := time.Now()
			v.UpdatedAt = &t
			v.Content = article.Content
			InMemoryArticle[i] = v

			return  nil
		}
	}
	return fmt.Errorf("article not found with id %s", article.Id)
}

func DeleteArticle(id string) error {
	for i, v := range InMemoryArticle {
		if v.Id == id && v.DeletedAt == nil {
			t := time.Now()
			InMemoryArticle[i].DeletedAt = &t
			return nil
		}
	}

	return fmt.Errorf("article not found with id %s", id)

}
