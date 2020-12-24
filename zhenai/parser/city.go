package parser

import (
	"goCrawler/engine"
	"regexp"
	"strings"
)

const cityRe = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func city(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, item := range matches {
		name := string(item[2])
		url := string(item[1])
		newUrl := strings.Replace(url, "http://", "https://", 1)
		result.Items = append(result.Items, name)
		result.Requests = append(result.Requests, engine.Request{
			Url:       newUrl,
			ParserFun: func(c []byte) engine.ParserResult {
				return profile(c, name)
			},
		})
	}
	return result
}
