package engine

import (
	"fetcher"
	"log"
)

type SimpleEngine struct {}

func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds{
		requests = append(requests, r)
	}

	for len(requests) > 0{
		r := requests[0]
		requests = requests[1:]

		ParserResult, err := worker(r)
		if err != nil{
			continue
		}

		requests = append(requests, ParserResult.Request...)

		for _, item := range ParserResult.Items{
			log.Printf("Got item %v", item)
		}
	}
}

func worker(r Request) (ParserResult, error){
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil{
		log.Printf("Fetcher : error" + "finding url %s: %v", r.Url, err)
		return ParserResult{}, err
	}
	return r.ParserFunc(body), err
}