package engine

import (
	"GoDemo/crawler/fetcher"
	"github.com/labstack/gommon/log"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}
	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]
		log.Printf("Fetching %s", r.Url)
		parserResult, err := worker(r)
		if err != nil {
			requests = append(requests, parserResult.Requests...)
		}
		for _, v := range parserResult.Items {
			log.Printf("Got item %v", v)
		}

	}
}
func worker(r Request) (ParseResult, error) {
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher:error Fetching url %s: %v", r.Url, err)
		return ParseResult{},err
	}
	parseResult := r.ParserFunc(body)
	return parseResult,nil
}
