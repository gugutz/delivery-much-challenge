package api

import (
	"delivery-much-api/api/recepies"
	"github.com/gin-gonic/gin"
)

func Start(router *gin.Engine) {
	recepies.ApplyRoutes(router) // apply api router
}
