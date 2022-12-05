package cards

import (
	"encoding/json"
	"fmt"
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

		if cardUrl != "" {
			card.URLs = append(card.URLs, cardUrl)
		}
		if imgUrl != "" {
			card.Photos = append(card.Photos, imgUrl)
		}

		card.buildHtml()
		return card

	}

	return nil
}
