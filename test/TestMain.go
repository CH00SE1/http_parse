package main

import (
	"httpParse/yellow"
	"sync"
)

/**
 * @title li5apuu7获取数据
 * @author xiongshao
 * @date 2022-06-27 17:08:58
 */

var wg sync.WaitGroup

func TE1(tag int) {
	for i := 2; i < 20; i++ {
		yellow.ExampleScrape(tag, i)
	}
}

func main() {
	TE1(24)
}
