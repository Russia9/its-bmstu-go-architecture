package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UpdateRequest struct {
	ID          string   `json:"id" binding:"required"`
	Title       string   `json:"title" binding:"required"`
	Content     string   `json:"content" binding:"required"`
	Attachments []string `json:"attachments"`
}

func (p *PostDelivery) Update(c *gin.Context) {
	var req UpdateRequest
	err := c.Bind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "wrong request format"})
		return
	}

	obj, err := p.uc.Update(c, req.ID, req.Title, req.Content, req.Attachments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": obj})
	return
}
