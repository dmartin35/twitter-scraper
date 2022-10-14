package twitterscraper

import ()

// unified card JSON object
type unifiedCard struct {
	Type string `json:"type"`

	/*
		Components       []string `json:"components"`
		ComponentObjects map[string]struct {
		} `json:"component_objects"`
		DestinationObjects map[string]struct {
		} `json:"destination_objects"`
	*/

	MediaEntities map[string]struct {
		IDStr         string `json:"id_str"`
		MediaURLHttps string `json:"media_url_https"`
		Type          string `json:"type"`
		URL           string `json:"url"`
	} `json:"media_entities"`
}

/*
{
  "type":"image_website",
  "component_objects":{
     "details_1":{
        "type":"details",
        "data":{
           "title":{
              "content":"Les maillots du SRFC aux ench\\u00e8res !",
              "is_rtl":false
           },
           "subtitle":{
              "content":"matchwornshirt.com",
              "is_rtl":false
           },
           "destination":"browser_1"
        }
     },
     "media_1":{
        "type":"media",
        "data":{
           "id":"3_1580862751999758337",
           "destination":"browser_1"
        }
     }
  },
  "destination_objects":{
     "browser_1":{
        "type":"browser",
        "data":{
           "url_data":{
              "url":"https://www.matchwornshirt.com/event/13-10-2022-dynamo-kyiv-stade-rennais-f-c",
              "vanity":"matchwornshirt.com"
           }
        }
     }
  },
  "components":[
     "media_1",
     "details_1"
  ],
  "media_entities":{
     "3_1580862751999758337":{
        "id":1580862751999758337,
        "id_str":"1580862751999758337",
        "indices":[
           0,
           0
        ],
        "media_url":"",
        "media_url_https":"https://pbs.twimg.com/media/FfBaTBOXkAEoyEj.jpg",
        "url":"",
        "display_url":"",
        "expanded_url":"",
        "type":"photo",
        "original_info":{
           "width":800,
           "height":800,
           "focus_rects":[
              {
                 "x":0,
                 "y":116,
                 "h":448,
                 "w":800
              },
              {
                 "x":0,
                 "y":0,
                 "h":800,
                 "w":800
              },
              {
                 "x":98,
                 "y":0,
                 "h":800,
                 "w":702
              },
              {
                 "x":380,
                 "y":0,
                 "h":800,
                 "w":400
              },
              {
                 "x":0,
                 "y":0,
                 "h":800,
                 "w":800
              }
           ]
        },
        "sizes":{
           "large":{
              "w":800,
              "h":800,
              "resize":"fit"
           },
           "thumb":{
              "w":150,
              "h":150,
              "resize":"crop"
           },
           "medium":{
              "w":800,
              "h":800,
              "resize":"fit"
           },
           "small":{
              "w":680,
              "h":680,
              "resize":"fit"
           }
        },
        "source_user_id":20447193,
        "source_user_id_str":"20447193",
        "media_key":"3_1580862751999758337",
        "ext":{
           "mediaColor":{
              "r":{
                 "ok":{
                    "palette":[
                       {
                          "rgb":{
                             "red":67,
                             "green":65,
                             "blue":66
                          },
                          "percentage":30.57
                       },
                       {
                          "rgb":{
                             "red":111,
                             "green":129,
                             "blue":47
                          },
                          "percentage":18.8
                       },
                       {
                          "rgb":{
                             "red":137,
                             "green":125,
                             "blue":115
                          },
                          "percentage":16.81
                       },
                       {
                          "rgb":{
                             "red":200,
                             "green":55,
                             "blue":49
                          },
                          "percentage":11.17
                       },
                       {
                          "rgb":{
                             "red":106,
                             "green":27,
                             "blue":33
                          },
                          "percentage":5.73
                       }
                    ]
                 }
              },
              "ttl":-1
           }
        }
     }
  }
}
*/

func (uc *unifiedCard) parse() *Card {

	return nil
}
