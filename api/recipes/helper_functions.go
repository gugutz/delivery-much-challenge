package recipes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"

	"errors"
	"strings"
)

func getGifFromGiphy(title string) (gifURL string, statusCode int, err error) {
	giphyURL := os.Getenv("GIPHY_URL")
	giphyKey := os.Getenv("GIPHY_APP_KEY")

	giphyURLTarget := fmt.Sprintf("%s/gifs/search?api_key=%s&q=%s&limit=1", giphyURL, giphyKey, title)

	// if the spaces in the URL target arent treated, the request returns 400
	giphyURLTarget = strings.ReplaceAll(giphyURLTarget, " ", "%20")

	// get the gif
	resp, err := http.Get(giphyURLTarget)
	if err != nil {
		log.Println("Error getting gif from Giphy: ", err)
	}
	// guard for the case when the spaces arent treated
	// returns in int that will be used by the router to respond with the correct HTTP statusCode
	if resp.StatusCode == 400 {
		errMsg := fmt.Sprintf("Bad request on Giphy API: %s", resp.Status)
		return "", 400, errors.New(errMsg)

	}
	if resp.StatusCode == 404 {
		errMsg := fmt.Sprintf("Error on call to Giphy API: %s", resp.Status)
		return "", 404, errors.New(errMsg)

	}
	if resp.StatusCode == 500 {
		errMsg := fmt.Sprintf("The Giphy API responded with error: %s", resp.Status)
		return "", 404, errors.New(errMsg)

	}

	// transforms http get response tu []bytes to unmarshall it
	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	// unmarshall the []byte response to json
	var giphyResponse GiphyResponse
	json.Unmarshal(body, &giphyResponse)

	var url string
	url = giphyResponse.Data[0].URL

	return url, 200, nil

}

//SortIngredients -
func sortIngredients(ingredients string) []string {
	explodeIngredients := strings.Split(ingredients, ",")

	for i := range explodeIngredients {
		explodeIngredients[i] = strings.TrimSpace(explodeIngredients[i])
	}
	sort.Strings(explodeIngredients)
	return explodeIngredients
}

func PrettyJson(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}
	return buffer.String(), nil
}
