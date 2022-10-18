package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/AbdulahadAbduqahhorov/gin/Article/docs"
	"github.com/AbdulahadAbduqahhorov/gin/Article/handlers"
	"github.com/AbdulahadAbduqahhorov/gin/Article/models"
	"github.com/AbdulahadAbduqahhorov/gin/Article/storage"
	"github.com/AbdulahadAbduqahhorov/gin/Article/storage/inmemory"
)

// @contact.name  API Article
// @contact.url   http://example.com
// @contact.email example@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	var stg storage.StorageI
	stg = inmemory.InMemory{
		Db: &inmemory.DB{},
	}
	r := gin.Default()
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"
	id := "7e2415ed-dd8f-40cd-a2a2-d16e253f1065"
	err:=stg.CreateAuthor(id, models.CreateAuthorModel{
		FirstName: "John",
		LastName:  "Doe",
	})
	if err != nil{
		panic(err)
	}
	idArticle := "f137fd5d-f8bf-46da-ad92-ac66dfadd634"
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
			panic(err)
		}
	}
	err = stg.CreateArticle(idArticle, models.CreateArticleModel{
		Content: models.Content{
			Title: "15",
			Body:  "News",
		},
		AuthorId: id,
	})
	if err != nil {
		panic(err)
	}
	err = stg.CreateArticle(idArticle, models.CreateArticleModel{
		Content: models.Content{
			Title: "16",
			Body:  "News",
		},
		AuthorId: id,
	})
	if err != nil {
		panic(err)
	}
	err = stg.CreateArticle(idArticle, models.CreateArticleModel{
		Content: models.Content{
			Title: "17",
			Body:  "News",
		},
		AuthorId: id,
	})
	if err != nil {
		panic(err)
	}

	err = stg.CreateArticle(idArticle, models.CreateArticleModel{
		Content: models.Content{
			Title: "18",
			Body:  "Sport",
		},
		AuthorId: id,
	})
	if err != nil {
		panic(err)
	}
	

	h := handlers.Handler{
		Stg: stg,
	}

	v1 := r.Group("v1")
	{
		v1.GET("/article", h.GetArticle)
		v1.POST("/article", h.CreateArticle)
		v1.PUT("/article", h.UpdateArticle)
		v1.DELETE("/article/:id", h.DeleteArticle)
		v1.GET("/article/:id", h.GetArticleById)

		v1.GET("/author", h.GetAuthor)
		v1.POST("/author", h.CreateAuthor)
		v1.PUT("/author", h.UpdateAuthor)
		v1.DELETE("/author/:id", h.DeleteAuthor)
		v1.GET("/author/:id", h.GetAuthorById)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
