package parser

import (
	"testing"
	"GoDemo/crawler/fetcher"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
)
var isFirst = true
func TestParserIpList(t *testing.T) {
	var count = 100;
	filePath := "IP.txt"
	//_, err := os.Stat(filePath)
	//if err == nil {
	//	os.Remove(filePath)
	//}
	for i := 2;i<= count;i++ {
		urlStr := "http://www.xicidaili.com/nt/"
		urlStr += strconv.Itoa(i)
		contents, err := fetcher.Fetch(urlStr)
		log.Print(string(contents))
		if err != nil {
			log.Printf("fail to fetch %s,the err is %s",urlStr, err)
		}
		result := ParserIpList(contents)
		WriteIp(result.Items,filePath) //todo should start a goroutine
	}
}
func WriteIp (items []interface{}, urlPath string) {
	file, err := os.OpenFile(urlPath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	if err != nil {
		log.Printf("fail to create file %s", urlPath)
	}


	defer file.Close()
	for _, v := range items {
		log.Printf("write %s to file", v)
		str := v.(string)
		str += "\n"
		file.WriteString(str)
	}

}