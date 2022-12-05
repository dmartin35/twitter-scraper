package cards

import (
	"fmt"
	"strings"
)

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

// buildHtml is a private function to generate the HTML code for the current Card struct
// and handle (optional/missing) image
func (card *Card) buildHtml() {
	var (
		imgTag string
	)

	if len(card.Photos) > 0 {
		imgTag = fmt.Sprintf(`<img src="%s"><br>`, card.Photos[0])
	}

	if len(card.URLs) > 0 {
		card.HTML = fmt.Sprintf(
			`<a href="%s">%s%s</a>`,
			card.URLs[0], imgTag, strings.Replace(card.Text, "\n", "<br>", -1))
	}

}
