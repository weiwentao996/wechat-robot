package main

import (
	"log"
	"sync"
	_ "wechatrobot/src/config"
	"wechatrobot/src/lib"
)

func main() {
	initCallback := make(chan bool, 1)
	go lib.Login(initCallback)
	select {
	case initResult := <-initCallback:
		if initResult {
			lib.Banner()
			block := sync.WaitGroup{}
			block.Add(1)
			block.Wait()
		}
	}
	log.Fatalln("程序出错！")
}
