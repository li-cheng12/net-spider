package parser

import (
	"GoDemo/crawler/engine"
	"regexp"
)

const IPParser = `<td>([\d.]*)</td>[\s\S]*?<td>([\d]*)</td>`

func ParserIpList (contents [] byte) engine.ParseResult {
	IpRegex := regexp.MustCompile(IPParser)
	IpMatches := IpRegex.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, v := range IpMatches {
		item := ""
		item += string(v[1])
		item += ":"
		item += string(v[2])
		result.Items = append(result.Items, item)
		result.Requests = append(result.Requests, engine.Request{
			Url:"",
			ParserFunc:engine.NilParse,
		})
	}
	return result
}
