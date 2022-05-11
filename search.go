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
		Title := string(data[i].Title)
		URL := string(data[i].URL)
		results = append(results, SearchResult{
			Title,
			URL})
	}
	return results
}

func convertType(result SearchResult) string{
	s := ""
	s += result.Title + "\n" + result.URL + "\n\n"
	return s
}

func createFile(path string) {
    var _, err = os.Stat(path)

    if os.IsNotExist(err) {
        var file, err = os.Create(path)
        if isError(err) {
            return
        }
        defer file.Close()
    }
}

func writeFile(path string, results []SearchResult) {
    var file, err = os.OpenFile(path, os.O_RDWR, 0644)
    if isError(err) {
        return
    }
    defer file.Close()

	for _, result := range results {
		resultString := convertType(result)
		_, err = file.WriteString(resultString)
		if isError(err) {
			return
		}
	}

    err = file.Sync()
    if isError(err) {
        return
    }
}

func isError(err error) bool {
    if err != nil {
        fmt.Println(err.Error())
    }

    return (err != nil)
}

func main() {
	// Here we get the google search term from the console arguments
	term := os.Args[1]
	data, _ := googlesearch.Search(nil, term)

	results := buildResults(data)

	var path = "results.txt"
	createFile(path)
	writeFile(path, results)
}