package router

import "github.com/gin-gonic/gin"

func Inicialize() {
	// Inicializa o Router utilizando as configurações padrões(Default) do gin
	r := gin.Default()
	// Define uma Rota
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// Estamos radando nossa api
	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
