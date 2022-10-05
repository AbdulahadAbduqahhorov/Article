package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/AbdulahadAbduqahhorov/gin/Article/docs"
	"github.com/AbdulahadAbduqahhorov/gin/Article/handlers"
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

	v1 := r.Group("v1")
	{
		v1.GET("/article", handlers.GetArticle)
		v1.POST("/article", handlers.CreateArticle)
		v1.PUT("/article", handlers.UpdateArticle)
		v1.DELETE("/article/:id", handlers.DeleteArticle)
		v1.GET("/article/:id", handlers.GetArticleById)

		v1.GET("/author", handlers.GetAuthor)
		v1.POST("/author", handlers.CreateAuthor)
		v1.PUT("/author", handlers.UpdateAuthor)
		v1.DELETE("/author/:id", handlers.DeleteAuthor)
		v1.GET("/author/:id", handlers.GetAuthorById)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run()
}
