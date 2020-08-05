package recepies

import (
	"log"

	"github.com/gin-gonic/gin"
)

// @Summary Login
// @Produce json
// @Param username query string true "userName"
// @Param password query string true "password"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func getListOfRecepies(ctx *gin.Context) {
	param := ctx.Params["i"]
	log.Println("param", param)

	var recipesResponse recipes.Response

	if err := ctx.BindJSON(&recipesResponse); err != nil {
		log.Println("Error on request body: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Something went wrong with the request",
			"status":  http.StatusBadRequest,
			"error":   err,
		})
		ctx.AbortWithStatus(400)
		return
	}

	// gin.H is a shortcut for map[string]interface{}
	ctx.JSON(http.StatusOK, gin.H{
		"keywords": "Login bem sucedido!",
		"recipes": gin.H{
			"user_id": &user.ID,
			// "access_token":  token,
			"access_token":  "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1OTMwNTQ0MjEsImlhdCI6IjIwMjAtMDYtMjRUMTY6MDc6MDEuNTA5ODY5LTAzOjAwIiwiaXNzIjoic3NvIiwic3ViIjoiMSJ9.GnFcnTqGMK7sf9dRMJ0hNQgxJZ1VwLdQFAfRKe-KD45",
			"refresh_token": "no-refresh-token-for-now",
		},
	})

}
