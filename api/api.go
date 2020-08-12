package api

import (
	"delivery-much-api/api/recipes"
	"github.com/gin-gonic/gin"
)

func Start(router *gin.Engine) {
	recipes.ApplyRoutes(router) // apply api router
}
