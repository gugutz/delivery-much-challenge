package recipes

type (
	// Recipe -
	Recipe struct {
		Title       string   `json:"title"`
		Ingredients []string `json:"ingredients"`
		Link        string   `json:"link"`
		GIF         string   `json:"gif"`
	}

	//Recipes -
	Recipes struct {
		KeyWords []string `json:"keywords"`
		Recipes  []Recipe `json:"recipes"`
	}

	PuppyRecipeResponse struct {
		Href    []interface{} `json:"href"`
		Results []struct {
			Title       string `json:"title"`
			Href        string `json:"href"`
			Ingredients string `json:"ingredients"`
		} `json:"results"`
	}

	// Giphy -
	GiphyResponse struct {
		Data []struct {
			URL string `json:"url"`
		} `json:"data"`
	}
)
