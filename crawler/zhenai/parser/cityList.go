package parser
import(
	"GoDemo/crawler/engine"
	"regexp"
)
const cityListRE = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`
func ParseCityList (contents [] byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRE)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, v:= range matches {
		result.Items = append(result.Items, string(v[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:string(v[1]),
			ParserFunc:ParsePersonList,
		})
	}
	return result
}
