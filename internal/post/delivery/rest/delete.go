package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (p *PostDelivery) Delete(c *gin.Context) {
	err := p.uc.Delete(c, c.Query("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"result": "ok"})
	return
}
