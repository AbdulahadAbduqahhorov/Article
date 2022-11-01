package inmemory

import (
	"errors"
	"strings"
	"time"

	"github.com/AbdulahadAbduqahhorov/gin/Article/models"
)

func (im InMemory) CreateAuthor(id string, author models.CreateAuthorModel) error {
	var response models.Author
	response.CreatedAt = time.Now()
	response.Id = id
	response.FirstName = author.FirstName
	response.LastName = author.LastName
	im.Db.InMemoryAuthor = append(im.Db.InMemoryAuthor, response)
	return nil
}

func (im InMemory) GetAuthor(limit, offset int, search string) (authors []models.Author) {

	count := 0

	for _, author := range im.Db.InMemoryAuthor {
		if author.DeletedAt == nil && (strings.Contains(author.FirstName, search) || strings.Contains(author.LastName, search)) {
			if count < offset {
				count++
			} else if limit > 0 {
				authors = append(authors, author)
				limit--

			}
		}
	}
	return
}

func (im InMemory) GetAuthorById(id string) (models.Author, error) {
	var author models.Author

	for _, v := range im.Db.InMemoryAuthor {
		if v.Id == id && v.DeletedAt == nil {
			return v, nil
		}
	}
	return author, errors.New("author not found")
}

func (im InMemory) UpdateAuthor(author models.UpdateAuthorModel) error {

	var response models.Author
	for i := range im.Db.InMemoryAuthor {
		if im.Db.InMemoryAuthor[i].Id == author.Id && im.Db.InMemoryAuthor[i].DeletedAt == nil {
			t := time.Now()
			response.UpdatedAt = &t
			response.CreatedAt = im.Db.InMemoryAuthor[i].CreatedAt
			response.FirstName = author.FirstName
			response.LastName = author.LastName
			response.Id = author.Id
			im.Db.InMemoryAuthor[i] = response

			return nil
		}
	}
	return errors.New("author not found with")
}

func (im InMemory) DeleteAuthor(id string) error {

	for i, v := range im.Db.InMemoryArticle {
		if v.Id == id && v.DeletedAt == nil {
			t := time.Now()
			im.Db.InMemoryArticle[i].DeletedAt = &t
			return nil
		}
	}

	return errors.New("article not found")

}
