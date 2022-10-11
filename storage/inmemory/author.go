package inmemory

import (
	"fmt"
	"time"

	"github.com/AbdulahadAbduqahhorov/gin/Article/models"
	"github.com/google/uuid"
)

var InMemoryAuthor []models.Author

func (im InMemory) CreateAuthor(author models.CreateAuthorModel) models.Author {
	var response models.Author
	t := time.Now()
	response.CreatedAt = &t
	id := uuid.New()
	response.Id = id.String()

	response.FirstName = author.FirstName
	response.LastName = author.LastName
	InMemoryAuthor = append(InMemoryAuthor, response)
	return response
}

func (im InMemory) GetAuthor() (authors []models.Author) {
	for _, author := range InMemoryAuthor {
		if author.DeletedAt == nil {
			authors = append(authors, author)
		}
	}
	return
}

func (im InMemory) GetAuthorById(id string) (models.Author, error) {
	var author models.Author

	for _, v := range InMemoryAuthor {
		if v.Id == id && v.DeletedAt == nil {
			return v, nil
		}
	}
	return author, fmt.Errorf("author not found with id %s", id)
}

func (im InMemory) UpdateAuthor(author models.UpdateAuthorModel) (models.Author, error) {

	var response models.Author
	for i := range InMemoryAuthor {
		if InMemoryAuthor[i].Id == author.Id && InMemoryAuthor[i].DeletedAt == nil {
			t := time.Now()
			response.UpdatedAt = &t
			response.CreatedAt = InMemoryAuthor[i].CreatedAt
			response.FirstName = author.FirstName
			response.LastName = author.LastName
			response.Id = author.Id
			InMemoryAuthor[i] = response

			return response, nil
		}
	}
	return response, fmt.Errorf("author not found with id %s", author.Id)
}

func (im InMemory) DeleteAuthor(id string) (models.Author, error) {
	author, err := im.GetAuthorById(id)
	if err != nil {
		return author, fmt.Errorf("author not found with id %s", id)
	}
	t := time.Now()
	author.DeletedAt = &t
	return author, nil

}
