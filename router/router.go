package router

import "github.com/gin-gonic/gin"

func Inicialize() {
	// Inicializa o Router
	router := gin.Default()

	// Inicializa o Router
	initializeRouter(router)

	// Estamos radando nossa api
	router.Run(":8080") // listen and serve on 0.0.0.0:8080
}
