package parser

import (
	"goCrawler/engine"
	"goCrawler/model"
	"regexp"
)

const profileRe = `<div class="m-btn purple">([^<]+)</div>`

func profile(content []byte, name string) engine.ParserResult {
	re := regexp.MustCompile(profileRe)
	matches := re.FindAllSubmatch(content, -1)
	result := engine.ParserResult{}
	for _, item := range matches {
		result.Items = append(result.Items, model.Profile{
			Name:          name,
			Status:        string(item[2]),
			Age:           string(item[3]),
			Constellation: string(item[4]),
			Height:        string(item[5]),
			Weight:        string(item[6]),
			Address:       string(item[7]),
			Income:        string(item[8]),
			Job:           string(item[9]),
			Education:     string(item[10]),
		})
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(item[1]),
			ParserFun: engine.NilParser,
		})
	}
	return result
}
