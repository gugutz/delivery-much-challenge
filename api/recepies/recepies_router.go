package recepies

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	recepiesRoutes := r.Group("/recipes")
	{
		recepiesRoutes.GET("/", getListOfRecepies)
	}
}
