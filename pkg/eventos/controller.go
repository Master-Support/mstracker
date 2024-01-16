package eventos

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(router *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	routes := router.Group("/api/v1/envio")
	routes.POST("", h.PostObjeto)
	routes.PUT("/atualizar", h.UpdateStatus)
	routes.GET("/objetos", h.GetObjetos)
	routes.GET("/objeto/:id", h.GetObjeto)
}
