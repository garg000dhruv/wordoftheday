package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type WordlistMetadata struct {
	Provider       string `json:"provider"`
	SourceLanguage string `json:"sourceLanguage"`
	Total          int    `json:"total"`
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
}

type WordMetadata struct {
	Id   string `json:"id"`
	Word string `json:"word"`
}

type Wordlist struct {
	Metadata WordlistMetadata `"json:metadata"`
	Results  []WordMetadata   `"json:results"`
}

func main() {
	var app_id = os.Args[1]
	var app_key = os.Args[2]
	var domainToFilter = os.Args[3]

	client := &http.Client{}

	wordlistReq, _ := http.NewRequest("GET", "https://od-api.oxforddictionaries.com:443/api/v1/wordlist/en/domains%3D"+domainToFilter, nil)

	wordlistReq.Header.Add("Accept", "application/json")
	wordlistReq.Header.Add("app_id", app_id)
	wordlistReq.Header.Add("app_key", app_key)

	wordlistResp, _ := client.Do(wordlistReq)

	defer wordlistResp.Body.Close()

	var filteredWordlist = new(Wordlist)
	filteredWordlistBody, _ := ioutil.ReadAll(wordlistResp.Body)
	filteredWordlistErr := json.Unmarshal(filteredWordlistBody, &filteredWordlist)
	if filteredWordlistErr != nil {
		panic(filteredWordlistErr)
	}

	fmt.Printf("Total words in domain '%s': %v\n", domainToFilter, len(filteredWordlist.Results))
}