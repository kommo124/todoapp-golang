package router

import (
	"todoapp/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() {
	r := gin.Default()
	r.Use(middlewares.CorsMiddleware())

}
