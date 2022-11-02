package cards

import (
	"encoding/json"
	"fmt"
	"strings"
)

// SummaryCardParser implements TwitterCardParser interface
type SummaryCardParser struct{}

func (p *SummaryCardParser) Parse(bindingValues map[string]interface{}) (card *Card) {

	var (
		title   string
		desc    string
		domain  string
		vanity  string
		cardUrl string
		imgUrl  string
	)

	if data, err := json.Marshal(bindingValues); err == nil {

		dotaccess(&title, data, "title.string_value")
		dotaccess(&desc, data, "description.string_value")
		dotaccess(&domain, data, "domain.string_value")
		dotaccess(&vanity, data, "vanity_url.string_value")
		dotaccess(&cardUrl, data, "card_url.string_value")
		dotaccess(&imgUrl, data, "thumbnail_image_original.image_value.url")

		card = &Card{}

		card.Text = fmt.Sprintf("%s\n%s\n%s\n", stringOr(vanity, domain), title, desc)
		card.HTML = card.Text

		card.URLs = append(card.URLs, cardUrl)
		card.Photos = append(card.Photos, imgUrl)

		if len(card.URLs) > 0 && len(card.Photos) > 0 {
			card.HTML = fmt.Sprintf(
				`<a href="%s"><img src="%s"><br>%s</a>`,
				card.URLs[0], card.Photos[0], strings.Replace(card.Text, "\n", "<br>", -1))
		}

		return card

	}

	return nil
}
