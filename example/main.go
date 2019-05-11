package main

import (
	"fmt"
	. "smart_assets/tool"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	// 初始化 mqtt client
	Init()
	wg.Add(1)
	// 初始化api
	err := InitApi()
	if err != nil {
		fmt.Println("初始化API失败")
		wg.Done()
	}
	wg.Wait()
}