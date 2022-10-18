package inmemory_test

import (
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/AbdulahadAbduqahhorov/gin/Article/models"
	"github.com/AbdulahadAbduqahhorov/gin/Article/storage/inmemory"
)

// test function

func InitData(stg *inmemory.InMemory) error {
	id := "1b705ff5-4038-48e7-b8f7-39f3a2ed3d64"
	err := stg.CreateAuthor(id, models.CreateAuthorModel{
		FirstName: "John",
		LastName:  "Doe",
	})
	if err != nil {
		return err
	}
	idArticle := "c470dd03-1e73-49fa-93a0-0c85fa16fc71"
	for i := 1; i <= 14; i++ {
		title := strconv.Itoa(i)
		err := stg.CreateArticle(idArticle, models.CreateArticleModel{
			Content: models.Content{
				Title: title,
				Body:  "Smth",
			},
			AuthorId: id,
		})
		if err != nil {
			return err
		}
	}
	return nil
}
func TestCreateArticle(t *testing.T) {
	var err error
	Im := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	AuthorID := "f6da1865-b8ae-42da-99f4-704b783da4cc"
	expectedError := errors.New("author not found")

	err = Im.CreateAuthor(AuthorID, models.CreateAuthorModel{
		FirstName: "Doe",
		LastName:  "John",
	})
	if err != nil {
		t.Fatalf("unknown error:%v", err)
	}
	content := models.Content{
		Title: "test",
		Body:  "test",
	}
	var tests = []struct {
		name       string
		id         string
		data       models.CreateArticleModel
		wantError  error
		wantResult models.GetArticleByIdModel
	}{
		{
			name: "success",
			id:   "2f395d6a-2202-4709-b669-9436d22c2ba8",
			data: models.CreateArticleModel{
				Content:  content,
				AuthorId: AuthorID,
			},
			wantError: nil,
			wantResult: models.GetArticleByIdModel{
				Content: content,
			},
		},
		{
			name: "failed",
			id:   "2f395d6a-2202-4709-b669-9436d22c2ba8",
			data: models.CreateArticleModel{
				Content:  content,
				AuthorId: "2f395d6a-2202-4709-b669-9436d22c2ba4",
			},
			wantError:  expectedError,
			wantResult: models.GetArticleByIdModel{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := Im.CreateArticle(tt.id, tt.data)

			if tt.wantError == nil {
				if err != nil {
					fmt.Errorf("Im.CreateArticle() got error %v", err)
				}
				article, err := Im.GetArticleById(tt.id)
				if err != nil {
					t.Errorf("got error: %v", err)
				}
				if article.Content != tt.wantResult.Content {
					t.Errorf("IM.AddArticle() expected: %v but got: %v", tt.wantResult.Content, article.Content)
				}
			} else {

				if err == nil {
					t.Errorf("Expected: %v but got nil", expectedError)
				} else if err.Error() != tt.wantError.Error() {
					t.Errorf("expected error: %v  but got error: %v", expectedError, err)
				}

			}

		})
	}
	t.Log("Test has been finished")
}
