package inmemory_test

import (
	"errors"
	"testing"

	"github.com/AbdulahadAbduqahhorov/gin/Article/models"
	"github.com/AbdulahadAbduqahhorov/gin/Article/storage/inmemory"
)

// test function
func TestCreateArticle(t *testing.T) {
	var err error

	id := "d53f9804-ef11-41a0-9aae-64fce0852cd1"
	Im := inmemory.InMemory{
		Db: &inmemory.DB{},
	}

	idAuthor := "f6da1865-b8ae-42da-99f4-704b783da4cc"
	err = Im.CreateAuthor(idAuthor, models.CreateAuthorModel{
		FirstName: "Doe",
		LastName:  "John",
	})
	if err != nil {
		t.Fatalf("unknown error:%v", err)
	}
	err = Im.CreateArticle(id, models.CreateArticleModel{
		Content: models.Content{
			Title: "hi",
			Body:  "hello",
		},
		AuthorId: idAuthor,
	})
	
	if err != nil{
		t.Errorf("got error: %v", err)
	}
	article,err:=Im.GetArticleById(id)
	if err != nil{
		t.Errorf("got error: %v", err)
	}

	if article.Title!="hi" && article.Body!="hello"{
		t.Errorf("mismatch between data")
	}






	err = Im.CreateArticle(id, models.CreateArticleModel{
		Content: models.Content{
			Title: "hi",
			Body:  "hello",
		},
		AuthorId: "3f1798c1-06f7-43f2-b45a-ed3dcd811dfe",
	})
	expectedError := errors.New("author not found")
	if err==nil{
		t.Errorf("Expected error but got nil")
	}else if err.Error() != expectedError.Error() {
		t.Errorf("expected error: %v  but got error: %v", expectedError, err)
	}


	
	
}
