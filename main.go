package main

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/readme-update-actions/structs"
)

// create new error for empty env
var errEnvEmpty = errors.New("getenv: Environment variable empty")

// Get env string from user through actions env
// if no env variables provided throw error
// requires key string
// returns string, error
func getEnvString(key string) (string, error) {
	env_string := os.Getenv(key)
	if env_string == "" {
		return env_string, errEnvEmpty
	}
	return env_string, nil
}

func main() {
	// get the rss list from the actions env
	rss_medium, _ := getEnvString("INPUT_RSS_LIST")

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
	var rss structs.RSS
	errXMLParse := xml.Unmarshal(responseBody, &rss)
	if errXMLParse != nil {
		log.Println("Error xml parse", errXMLParse)
	}

	fmt.Println(rss.Channel.Title)
}
