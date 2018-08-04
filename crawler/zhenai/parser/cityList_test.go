package parser

import (
	"testing"
	"io/ioutil"
)

func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("./test")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", contents)
	result := ParseCityList(contents)
	const resultSize = 470
	if len(result.Requests) != resultSize {
		t.Errorf("result should hava %d " + "requests; but had %d", resultSize, len(result.Requests))
	}
}
