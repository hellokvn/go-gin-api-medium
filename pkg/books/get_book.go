package books

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hellokvn/go-api-medium/pkg/common/models"
)

func (h handler) GetBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	if result := h.DB.First(&book, id); result.Error != nil {
		c.Status(http.StatusNotFound)

		return
	}

	c.JSON(http.StatusOK, &book)
}
