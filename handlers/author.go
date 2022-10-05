package handlers

import (
	"net/http"

	"github.com/AbdulahadAbduqahhorov/gin/Article/models"
	"github.com/AbdulahadAbduqahhorov/gin/Article/storage"
	"github.com/gin-gonic/gin"
)

// CreateAuthor godoc
// @Summary     Create author
// @Description create author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     models.CreateAuthorModel true "author body"
// @Success     201     {object} models.JSONResult{data=models.Author}
// @Failure     400     {object} models.JSONErrorResult
// @Router      /v1/author [post]
func CreateAuthor(c *gin.Context) {
	var author models.CreateAuthorModel

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}
	response := storage.CreateAuthor(author)

	c.JSON(http.StatusCreated, models.JSONResult{
		Message: "Author created",
		Data:    response,
	})
}

// GetAuthorList godoc
// @Summary     List Author
// @Description get author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Success     200 {object} models.JSONResult{data=[]models.Author}
// @Router      /v1/author [get]
func GetAuthor(c *gin.Context) {
	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Author | GetList",
		Data:    storage.GetAuthor(),
	})
}

// GetAuthorById godoc
// @Summary     Get author by id
// @Description get an author by id
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       id  path     string true "author id"
// @Success     200 {object} models.JSONResult{data=models.Author}
// @Failure     400 {object} models.JSONErrorResult
// @Router      /v1/author/{id} [get]
func GetAuthorById(c *gin.Context) {
	id := c.Param("id")

	res, err := storage.GetAuthorById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Author | GetById",
		Data:    res,
	})
}

// UpdateAuthor godoc
// @Summary     Update author
// @Description update author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author body     models.UpdateAuthorModel true "author body"
// @Success     200     {object} models.JSONResult{data=models.Author}
// @Failure     400     {object} models.JSONErrorResult
// @Router      /v1/author [put]
func UpdateAuthor(c *gin.Context) {
	var author models.UpdateAuthorModel

	if err := c.ShouldBindJSON(&author); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	res, err := storage.UpdateAuthor(author)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Author | Update",
		Data:    res,
	})
}

// DeleteAuthor godoc
// @Summary     Delete author
// @Description delete an author
// @Tags        authors
// @Produce     json
// @Param       id  path     string true "author id"
// @Success     200 {object} models.JSONResult{data=models.Author}
// @Failure     400 {object} models.JSONErrorResult
// @Router      /v1/author/{id} [delete]
func DeleteAuthor(c *gin.Context) {
	id := c.Param("id")

	res, err := storage.DeleteAuthor(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResult{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResult{
		Message: "Author Deleted",
		Data:    res,
	})
}
