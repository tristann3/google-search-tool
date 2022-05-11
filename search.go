package main

import (
	"fmt"
	"github.com/rocketlaunchr/google-search"
	"os"
)

func main() {
	fmt.Println(googlesearch.Search(nil, "cars for sale in Toronto, Canada"))

	term := os.Args[1]
	fmt.Println(term)

}