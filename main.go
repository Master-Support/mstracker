package main

import (
	"example.com/mstracker/config"
	"example.com/mstracker/pkg/eventos"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)


func main() {
	config.CarregaVariaveisDeAmbiente()

	r := gin.Default()
	dbHundle := config.ConectaDB()
	
	eventos.RegisterRoutes(r, dbHundle)


	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"port": "Conectou mano",
		})
	})
	r.Run()
}
