package utils

import (
	"os"
	"github.com/labstack/gommon/log"
	"bufio"
	"io"
	"strings"
)

var IP = []string{}
func LoadingIP (filePath string) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0666)
	if err != nil {
		log.Printf("fail to open %s", filePath)
	}
	defer file.Close()
	buf := bufio.NewReader(file)
	count := 0
	for {
		line, err := buf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				log.Print("file has been to the end")
				break
			} else {
				log.Printf("it has a file error")
			}

		}
		line = strings.TrimSpace(line)
		IP[count] = line
		count ++
	}


}

