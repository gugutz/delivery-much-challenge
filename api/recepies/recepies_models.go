package recepies

type (
	Response struct {
		Keywords []string `json:"keywords"`
		Recipes  []struct {
			Title       string   `json:"title"`
			Ingredients []string `json:"ingredients"`
			Link        string   `json:"link"`
			Gif         string   `json:"gif"`
		} `json:"recipes"`
	}
)
