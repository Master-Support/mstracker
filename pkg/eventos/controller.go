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
    routes := router.Group("/objetos-correios")
    routes.POST("", h.PostObjeto)
    routes.GET("/objetos", h.GetObjetos)
    routes.GET("/:id", h.GetObjeto)
}
