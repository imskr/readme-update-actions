package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	medium "github.com/readme-update-actions/pkg/structs"
	helpers "github.com/readme-update-actions/pkg/utils"
)

func main() {
	// get the rss list from the actions env
	rss_medium, _ := helpers.GetEnvString("INPUT_RSS_LIST")

	// get the number of posts or stories to commit
	max_post, _ := helpers.GetEnvInteger("INPUT_MAX_POST")

	// if max_post not in env var set default to 3
	if max_post == 0 {
		max_post = 3
	}

	// get readme path from the actions env
	readme_path, _ := helpers.GetEnvString("INPUT_README_PATH")

	// if path not provided default to root readme
	if readme_path == "" {
		readme_path = "./README.md"
	}

	// get medium.com rss feed
	mediumResponse, err := http.Get(rss_medium)
	if err != nil {
		log.Println("Error making request to medium", err)
	}

	defer mediumResponse.Body.Close()

	responseBody, err := ioutil.ReadAll(mediumResponse.Body)
	if err != nil {
		log.Println("Error reading response body", err)
	}

	// use RSS structs
	var rss medium.RSS
	errXMLParse := xml.Unmarshal(responseBody, &rss)
	if errXMLParse != nil {
		log.Println("Error xml parse", errXMLParse)
	}

	// store the posts
	var items []string

	// get the posts
	// format it according to readme links format
	for i := 0; i < max_post; i++ {
		item := fmt.Sprintf("- [%s](%s)\n", rss.Channel.Item[i].Title, rss.Channel.Item[i].Link)
		items = append(items, item)
	}
	result_post := fmt.Sprintf("<!-- BLOG-LIST-START -->"+"\n%s", strings.Join(items, "\n"))

	// find readme and replace with our result
	helpers.ReplaceFile(readme_path, strings.TrimSuffix(result_post, "\n"))
}
