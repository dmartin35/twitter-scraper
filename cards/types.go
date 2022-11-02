package cards

type (
	// Card type is the public struct for embedded card into tweets.
	Card struct {
		//ID     string   `json:"id"`
		Text   string   `json:"text"`
		HTML   string   `json:"html"`
		Photos []string `json:"photos"`
		URLs   []string `json:"urls"`
	}

	// TwitterCard defines the interface for parsing different twitter cards (summary, unified, ...)
	TwitterCardParser interface {
		Parse(bindingValues map[string]interface{}) *Card
	}
)
