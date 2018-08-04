package parser

import (
	"GoDemo/crawler/engine"
	"regexp"
)
const personPatser = `<a href="(http://album.zhenai.com/u/[0-9]*)" target="_blank">([^<]+)</a>`
const urlParser = `<a href="(http://www.zhenai.com/zhenghun/[^"]+)"`
//目前先爬第一页
func ParsePersonList(contents []byte) engine.ParseResult {
	persionRe := regexp.MustCompile(personPatser)
	matches := persionRe.FindAllSubmatch(contents, -1)
	parserResult := engine.ParseResult{}
	for _, v := range matches {
		name := string(v[2])
		parserResult.Items = append(parserResult.Items, string(v[2]))
		parserResult.Requests = append(parserResult.Requests, engine.Request{
			Url:string(v[1]),
			ParserFunc: func(bytes []byte) engine.ParseResult {
				return ParseProfile(bytes, name)
			},
		})
	}

	urlRe := regexp.MustCompile(urlParser)
	urlMatches := urlRe.FindAllSubmatch(contents, -1)
	for _, v := range urlMatches {
		parserResult.Requests = append(parserResult.Requests, engine.Request{
			Url : string(v[1]),
			ParserFunc: ParsePersonList,
		})
	}
	return parserResult

}
