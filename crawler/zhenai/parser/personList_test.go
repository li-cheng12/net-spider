package parser

import (
	"testing"
	"fmt"
	"io/ioutil"
)

func TestParsePersonList(t *testing.T) {
	//contents, err := fetcher.Fetch("http://www.zhenai.com/zhenghun/aba")
	//ioutil.WriteFile("personTest.txt", contents, os.ModeAppend)
	contents, err := ioutil.ReadFile("personTest.txt")
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", contents)
	result := ParsePersonList(contents)
	for _, r := range result.Items {
		fmt.Printf("persons' name is %s\n", r)
	}
}
