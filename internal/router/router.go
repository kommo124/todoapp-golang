package router

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"todoapp/internal/middlewares"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.CorsMiddleware())

	return r
}
