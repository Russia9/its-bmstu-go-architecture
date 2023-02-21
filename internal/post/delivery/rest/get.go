package rest

import (
	"errors"
	"github.com/Russia9/its-bmstu-go-architecture/pkg/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PostDelivery) Get(c *gin.Context) {
	get, err := p.uc.Get(c, c.GetString("id"))
	if errors.Is(err, domain.ErrPostNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "post not found",
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"result": get,
	})
	return
}
