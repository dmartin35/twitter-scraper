package twitterscraper

import (
	"encoding/json"
	"fmt"
	"strings"
)

// unified card JSON object
type unifiedCard struct {
	Type string `json:"type"`

	Components       []string `json:"components"`
	ComponentObjects map[string]struct {
		Type string                 `json:"type"`
		Data map[string]interface{} `json:"data"`
	} `json:"component_objects"`
	DestinationObjects map[string]struct {
		Type string `json:"type"`
		Data struct {
			UrlData struct {
				URL    string `json:"url"`
				Vanity string `json:"vanity"`
			} `json:"url_data"`
		} `json:"data"`
	} `json:"destination_objects"`

	MediaEntities map[string]struct {
		IDStr         string `json:"id_str"`
		MediaURLHttps string `json:"media_url_https"`
		Type          string `json:"type"`
		URL           string `json:"url"`
	} `json:"media_entities"`
}

// unified card's details JSON object
type details struct {
	Title struct {
		Content string `json:"content"`
	} `json:"title"`
	Subtitle struct {
		Content string `json:"content"`
	} `json:"subtitle"`
	Destination string `json:"destination"`
}

func (uc *unifiedCard) parseCard() *Card {
	// regardless multiple keys holding structures,
	// there is only on browser link, detail and media
	// for both image & video cards

	knownTypes := []string{"image_website", "video_website"}
	if !stringInSlice(uc.Type, knownTypes) {
		return nil
	}

	c := &Card{}

	for _, component := range uc.ComponentObjects {
		if component.Type == "details" {

			var d details
			jsondetails, err := json.Marshal(component.Data)
			if err != nil {
				continue
			}
			if err := json.Unmarshal(jsondetails, &d); err != nil {
				continue
			}

			c.Text = fmt.Sprintf("%s\n%s", d.Subtitle.Content, d.Title.Content)
			c.HTML = c.Text
		}
	}

	for _, dest := range uc.DestinationObjects {
		switch dest.Type {
		case "browser", "browser_with_docked_media":
			c.URLs = append(c.URLs, dest.Data.UrlData.URL)
		}
	}

	for _, media := range uc.MediaEntities {
		switch media.Type {
		case "photo", "video":
			c.Photos = append(c.Photos, media.MediaURLHttps)
		}
	}

	if len(c.URLs) > 0 && len(c.Photos) > 0 {
		c.HTML = fmt.Sprintf(
			`<a href="%s"><img src="%s"><br>%s</a>`,
			c.URLs[0], c.Photos[0], strings.Replace(c.Text, "\n", "<br>", 1))
	}

	return c
}
