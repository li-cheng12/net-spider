package engine

type ParseResult struct {
	Requests [] Request
	Items []interface{}
}

type Request struct {
 	Url string
 	ParserFunc func([] byte) ParseResult
 }

 func NilParse([] byte) ParseResult {
	 return ParseResult{}
 }
