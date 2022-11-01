package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/AbdulahadAbduqahhorov/gin/Article/docs"
	"github.com/AbdulahadAbduqahhorov/gin/Article/handlers"
	"github.com/AbdulahadAbduqahhorov/gin/Article/storage"
	"github.com/AbdulahadAbduqahhorov/gin/Article/storage/postgres"
)

// @contact.name  API Article
// @contact.url   http://example.com
// @contact.email example@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	r := gin.Default()
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server Petstore server."
	docs.SwaggerInfo.Version = "1.0"

	var stg storage.StorageI
	var err error
	stg, err = postgres.InitDb("user=abdulahad password=10082018 dbname=uacademy sslmode=disable")
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
