package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

func main() {
	var app_id = os.Args[1]
	var app_key = os.Args[2]
	var domainToFilter = os.Args[3]

	// The default number generator is deterministic, so it'll
	// produce the same sequence of numbers each time by default.
	// To produce varying sequences, give it a seed that changes.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	client := &http.Client{}

	// Retrieve a list of words for the given domain.
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

	// Retrieve information for a randomly selected word from the domain-filtered list.
	var randomSelectedWord = filteredWordlist.Results[r1.Intn(len(filteredWordlist.Results))].Word

	wordDefReq, _ := http.NewRequest("GET", "https://od-api.oxforddictionaries.com:443/api/v1/entries/en/"+randomSelectedWord, nil)

	wordDefReq.Header.Add("Accept", "application/json")
	wordDefReq.Header.Add("app_id", app_id)
	wordDefReq.Header.Add("app_key", app_key)

	wordDefResp, _ := client.Do(wordDefReq)

	defer wordDefResp.Body.Close()

	var selectedWordMetadata = new(Wordlist)
	wordDefBody, _ := ioutil.ReadAll(wordDefResp.Body)
	selectedWordMetadataErr := json.Unmarshal(wordDefBody, &selectedWordMetadata)
	if selectedWordMetadataErr != nil {
		panic(selectedWordMetadataErr)
	}

	fmt.Printf("Word of the Day in domain '%s': %s\n", domainToFilter, selectedWordMetadata.Results[0].Word)

	// Print the different definitions for the selected word under every lexical category.
	for _, lexicalEntry := range selectedWordMetadata.Results[0].LexicalEntries {
		fmt.Printf("[%s]\n", lexicalEntry.LexicalCategory)
		for _, entry := range lexicalEntry.Entries {
			for index, sense := range entry.Senses {
				fmt.Printf("\tDefinition %v: %v\n", index+1, sense.Definitions)
			}
		}
	}
}
