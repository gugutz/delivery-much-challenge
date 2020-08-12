package recipes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"strings"
)

// @Summary GetRecipes
// @Produce json
// @Param i query string true "i"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /auth [post]
func getListOfRecipes(ctx *gin.Context) {
	// get the querystring param with the ingredients
	params := ctx.Query("i")

	// get the recipes
	recipePuppyURL := fmt.Sprintf("http://www.recipepuppy.com/api/?i=%s", params)
	resp, err := http.Get(recipePuppyURL)
	if err != nil {
		log.Println("Error getting recipes from Recipe Puppy: ", err)
	}

	// transforms http get response tu []bytes to unmarshall it
	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	// unmarshall to json
	var puppyRecipes PuppyRecipeResponse
	json.Unmarshal(body, &puppyRecipes)

	// save recipes into required response struct
	var recipes []Recipe
	recipes = make([]Recipe, len(puppyRecipes.Results))
	for index, puppyRecipe := range puppyRecipes.Results {
		recipes[index].Title = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(puppyRecipe.Title, "\r", ""), "\t", ""), "\n", "")
		recipes[index].Ingredients = sortIngredients(puppyRecipe.Ingredients)
		recipes[index].Link = puppyRecipe.Href

		// get the gif url to save for this recipe
		gifURL, statusCode, err := getGifFromGiphy(recipes[index].Title)
		if err != nil {
			if statusCode == 400 {
				errMsg := fmt.Sprintf("Error getting gif from Giphy: %s", err)
				log.Println(errMsg)
				ctx.JSON(http.StatusBadRequest, gin.H{
					"message": errMsg,
					"status":  http.StatusBadRequest,
					"error":   err,
				})
				ctx.AbortWithStatus(400)
				return
			}
		}
		recipes[index].GIF = gifURL
	}

	keywords := strings.Split(params, ",")

	ctx.JSON(http.StatusOK, gin.H{
		"keywords": keywords,
		"recipes":  recipes,
	})

}
