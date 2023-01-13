package lib

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func Banner() {
	file, err := os.OpenFile("./banner", os.O_RDONLY, 0444)
	defer file.Close()
	if err != nil {
		log.Println("读取banner失败！")
	}
	strBytes := make([]byte, 1024*5)
	buf := bufio.NewReader(file)
	n, err := buf.Read(strBytes)
	if err != nil && err != io.EOF {
		log.Println("读取banner失败！")
	}
	fmt.Println(string(strBytes[:n]))
}
