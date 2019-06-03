package main

import (
	"fmt"
	"log"
	. "smart_assets/tool"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// 初始化 mqtt client
	log.Println("连接mqtt broker")
	Init(CLIENTID)
	wg.Add(1)
	// 初始化api
	log.Println("注册api")
	err := InitApi()
	if err != nil {
		fmt.Println("初始化API失败")
		wg.Done()
	}
	wg.Wait()
}