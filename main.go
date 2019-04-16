package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
)

const profileName = ""

func main() {
	resp, err := http.Get(fmt.Sprintf("https://instagram.com/%s", profileName))

	if err != nil {
		log.Println(err.Error())
		return
	}

	defer resp.Body.Close()

	z := html.NewTokenizer(resp.Body)
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			// End of the document, we're done
			return
		case tt == html.StartTagToken:
			t := z.Token()

			// Look for all script tags
			isScript := t.Data == "script"
			if isScript {
				z.Next()
				t = z.Token()
				// Instagram puts the post json in a window._sharedData object
				if !strings.HasPrefix(t.Data, "window._sharedData = ") {
					continue
				}

				if strings.HasPrefix(t.Data, "window._sharedData = ") {
					log.Println("Found sharedData")
				}

				// Split to grab just the json piece
				sArray := strings.Split(t.Data, " = ")
				if len(sArray) < 1 {
					log.Println("Something went wrong!")
					continue
				}

				jsonMap := sArray[1]

				// Remove the trailing semicolon to make valid json
				jsonMap = strings.TrimSuffix(jsonMap, ";")

				// Unmarshal json contents
				var contents InstagramDto
				err := json.Unmarshal([]byte(jsonMap), &contents)
				if err != nil {
					continue
				}

				responseArray := make([]string, 0)
				// Parse this ugly Graphql object as though it is json :S
				for _, v := range contents.EntryData.ProfilePage[0].Graphql.User.EdgeOwnerToTimelineMedia.Edges {
					responseArray = append(responseArray, v.Node.Shortcode)
				}

				log.Println("Shortcodes: ", responseArray)
				return
			}
		}
	}
}
