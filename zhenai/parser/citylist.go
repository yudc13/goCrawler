package parser

import (
	"goCrawler/engine"
	"regexp"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func CityList(contents []byte) engine.ParserResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParserResult{}
	for _, match := range matches {
		// 	保存城市名称
		result.Items = append(result.Items, string(match[2]))
		// 保存接下来需要处理的url已经对应的解析器
		result.Requests = append(result.Requests, engine.Request{
			Url:       string(match[1]),
			ParserFun: engine.NilParser,
		})
	}
	return result
}
