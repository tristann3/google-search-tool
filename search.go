package main

import (
	"fmt"
	"github.com/rocketlaunchr/google-search"
	"os"
)

type SearchResult struct {
	Title string
	URL string
}

func buildResults(data []googlesearch.Result) []SearchResult{
	numResults := len(data)
	var results []SearchResult
	for i := 0; i < numResults; i++ {
		results = append(results, SearchResult{
			data[i].Title,
			data[i].URL})
	}
	return results
}

func main() {
	// Here we get the google search term from the console arguments
	term := os.Args[1]
	fmt.Println(term)

	data, _ := googlesearch.Search(nil, term)

	results := buildResults(data)
	fmt.Println(results[1].URL)


}