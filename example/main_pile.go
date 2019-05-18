package main

import (
	"fmt"
	"sync"

	. "smart_assets/tool"

	. "smart_assets/pile"
)

func main() {
	var wg sync.WaitGroup
	// 初始化 mqtt client
	Init()
	wg.Add(1)
	// 初始化api
	err := InitPileApi()
	if err != nil {
		fmt.Println("初始化API失败")
		wg.Done()
	}
	wg.Wait()
}