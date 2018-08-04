package main

import (
	_"golang.org/x/text"
	_"golang.org/x/net/html"
	"GoDemo/crawler/engine"
	"GoDemo/crawler/schedule"
	"GoDemo/crawler/zhenai/parser"
)


func main() {
	//engine.Run(engine.Request{
	//	Url:"http://www.zhenai.com/zhenghun",
	//	ParserFunc:parser.ParseCityList,
	//})
	concurEngin := engine.ConcurrentEngin{
		WorkCount:10,
		Schedule: &schedule.SimpleSchedule{},
	}
	//request := engine.Request{
	//	Url : "http://www.zhenai.com/zhenghun",
	//	ParserFunc:parser.ParseCityList,
	//}
	request := engine.Request{
		Url:"http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc:parser.ParsePersonList,
	}
	concurEngin.Run(request)
}
