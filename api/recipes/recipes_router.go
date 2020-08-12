package recipes

import (
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.Engine) {
	recipesRoutes := r.Group("/recipes")
	{
		recipesRoutes.GET("/", getListOfRecipes)
	}
}
