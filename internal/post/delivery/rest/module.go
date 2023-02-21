package rest

import (
	"github.com/Russia9/its-bmstu-go-architecture/pkg/domain"
	"github.com/gin-gonic/gin"
)

type PostDelivery struct {
	uc domain.PostUsecase
}

func NewPostDelivery(uc domain.PostUsecase, g *gin.RouterGroup) {
	module := &PostDelivery{uc}

	g.GET("/get", module.Get)
	g.POST("/create", module.Create)
	g.POST("/update", module.Update)
	g.DELETE("/delete", module.Delete)
}
