package engine

import (
	"fmt"
	"goCrawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request
	requests = append(requests, seeds...)
	for len(requests) > 0 {
		current := requests[0]
		requests = requests[1:]
		log.Printf("fetching %s \n", current.Url)
		contents, err := fetcher.Fetch(current.Url)
		if err != nil {
			log.Printf("fetcher fetch err url: %s err: %v", current.Url, err)
			continue
		}
		parserResult := current.ParserFun(contents)
		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			fmt.Printf("fetch item %v \n", item)
		}
	}
}
