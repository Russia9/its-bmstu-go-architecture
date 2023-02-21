package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CreateRequest struct {
	Title       string   `json:"title" binding:"required"`
	Content     string   `json:"content" binding:"required"`
	Attachments []string `json:"attachments"`
}

func (p *PostDelivery) Create(c *gin.Context) {
	var req CreateRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong request format"})
		return
	}

	obj, err := p.uc.Create(c, req.Title, req.Content, req.Attachments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": obj})
	return
}
