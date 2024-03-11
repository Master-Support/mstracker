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

	router.Use(corsMiddleware())

	routes := router.Group("/api/v1/envio")
	routes.POST("", h.PostObjeto)
	routes.PUT("/atualizar", h.UpdateStatus)
	routes.GET("/objetos", h.GetObjetos)
	routes.GET("/objeto/:id", h.GetObjeto)
}

func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, CREATE,OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
			return
		} else {
			c.Next()
		}

	}
}
