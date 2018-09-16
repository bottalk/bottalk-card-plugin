package main

import (
	"encoding/json"
	"net/http"

	bottalk "github.com/bottalk/go-plugin"
)

type btRequest struct {
	Token  string            `json:"token"`
	UserID string            `json:"user"`
	Vars   map[string]string `json:"vars"`
}

// Card represents alexa card object
type Card struct {
	Type  string           `json:"type"`
	Title string           `json:"title"`
	Text  string           `json:"text"`
	Image ImageDescription `json:"image"`
}

// ImageDescription represents small and large images
type ImageDescription struct {
	SmallImage string `json:"smallImageUrl"`
	LargeImage string `json:"largeImageUrl"`
}

func errorResponse(message string) string {
	return "{\"result\": \"fail\",\"message\":\"" + message + "\"}"
}

func main() {

	plugin := bottalk.NewPlugin()
	plugin.Name = "Simple Card Plugin"
	plugin.Description = "This plugin outputs simple alexa card"

	plugin.Actions = map[string]bottalk.Action{"card": bottalk.Action{
		Name:        "card",
		Description: "This action output simple card",
		Endpoint:    "/card",
		Action: func(r *http.Request) string {

			var BTR btRequest
			decoder := json.NewDecoder(r.Body)

			err := decoder.Decode(&BTR)
			if err != nil {
				return errorResponse(err.Error())
			}

			card := Card{
				Type:  "Standard",
				Title: BTR.Vars["title"],
				Text:  BTR.Vars["text"],
				Image: ImageDescription{
					SmallImage: BTR.Vars["image"],
					LargeImage: BTR.Vars["image"],
				},
			}

			b, err := json.Marshal(card)

			return "{\"result\": \"ok\",\"vars\":{\"card\":" + string(b) + "}}"
		},
		Params: map[string]string{
			"text":  "Text to display in a card",
			"title": "Title of the card",
			"image": "Image url",
		},
	}}

	plugin.Run(":9081")
}
