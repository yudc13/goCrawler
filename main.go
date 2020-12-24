package main

import (
	"goCrawler/engine"
	"goCrawler/zhenai/parser"
)

func main() {
	url := "http://www.zhenai.com/zhenghun"
	engine.Run(engine.Request{
		Url:       url,
		ParserFun: parser.CityList,
	})
}
