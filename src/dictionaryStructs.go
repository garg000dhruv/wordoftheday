package main

// Additional information provided by OUP for any Wordlist.
type WordlistMetadata struct {
	Provider       string `json:"provider"`
	SourceLanguage string `json:"sourceLanguage"`
	Limit          int    `json:"limit"`
	Offset         int    `json:"offset"`
	Total          int    `json:"total"`
}

// Entry containing everything pertaining to a particular word in OUP.
type WordMetadata struct {
	Id             string         `json:"id"`
	Language       string         `json:"language"`
	Type           string         `json:"type"`
	Word           string         `json:"word"`
	LexicalEntries []LexicalEntry `json:"lexicalEntries"`
}

// A wordlist is a group of words which share a common feature (e.g. all nouns,
// all words related to cricket, all words labelled as archaic).
type Wordlist struct {
	Metadata WordlistMetadata `"json:metadata"`
	Results  []WordMetadata   `"json:results"`
}

// A particular meaning of a word.
// Example: the word ‘mouse’ has the following senses:
//  - A small rodent that typically has a pointed snout, relatively large ears and eyes, and a long tail.
//  - A small handheld device which is moved across a mat or flat surface to move the cursor on a computer screen.
type Sense struct {
	Id          string    `json:"id"`
	Definitions []string  `json:"definitions"`
	Domains     []string  `json:"domains"`
	Examples    []Example `json:"examples"`
}

type Example struct {
	Text string `json:"text"`
}

// A complete account of a particular word.
// This can include a word’s senses, definitions, translations, origin, and any phrases featuring the word.
type Entry struct {
	Etymologies []string `json:"etymologies"`
	Senses      []Sense  `json:"senses"`
}

// A grouping of various senses in a specific language, and a lexical category that relates to a word.
type LexicalEntry struct {
	Language        string  `json:"language"`
	LexicalCategory string  `json:"lexicalCategory"`
	Text            string  `json:"text"`
	Entries         []Entry `json:"entries"`
}
