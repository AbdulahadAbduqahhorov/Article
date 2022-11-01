package inmemory

import (
	"errors"
	"strings"
	"time"

	"github.com/AbdulahadAbduqahhorov/gin/Article/models"
)

func (im InMemory) CreateArticle(id string, article models.CreateArticleModel) error {
	var response models.Article
	_, err := im.GetAuthorById(article.AuthorId)
	if err != nil {
		return err
	}
	t := time.Now()
	response.CreatedAt = t
	response.Id = id
	response.Content = article.Content
	response.AuthorId = article.AuthorId
	im.Db.InMemoryArticle = append(im.Db.InMemoryArticle, response)

	return nil

}

func (im InMemory) GetArticle(limit, offset int, search string) (articles []models.Article) {
	count:=0

	for _, article := range im.Db.InMemoryArticle {
		if article.DeletedAt == nil && (strings.Contains(article.Title, search) || strings.Contains(article.Body, search)) {
			if count < offset {
				count++
			}else if limit>0{
				articles = append(articles, article)
				limit--
				
			}
		}
	}
	return
}

func (im InMemory) GetArticleById(id string) (models.GetArticleByIdModel, error) {
	var article models.GetArticleByIdModel

	for _, v := range im.Db.InMemoryArticle {
		if v.Id == id && v.DeletedAt == nil {
			author, err := im.GetAuthorById(v.AuthorId)
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
	return article, errors.New("article not found")
}

func (im InMemory) UpdateArticle(article models.UpdateArticleModel) error {

	for i, v := range im.Db.InMemoryArticle {
		if im.Db.InMemoryArticle[i].Id == article.Id && im.Db.InMemoryArticle[i].DeletedAt == nil {
			t := time.Now()
			v.UpdatedAt = &t
			v.Content = article.Content
			im.Db.InMemoryArticle[i] = v

			return nil
		}
	}
	return errors.New("article not found")
}

func (im InMemory) DeleteArticle(id string) error {
	for i, v := range im.Db.InMemoryArticle {
		if v.Id == id && v.DeletedAt == nil {
			t := time.Now()
			im.Db.InMemoryArticle[i].DeletedAt = &t
			return nil
		}
	}

	return errors.New("article not found")

}
