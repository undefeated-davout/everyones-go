package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	queue := make(chan string)
	for i := 0; i < 2; i++ { // 2つのゴルーチン（ワーカー）を生成
		wg.Add(1)
		go fetchURL(queue, i)
	}

	queue <- "https://www.example.com"
	queue <- "https://www.example.net"
	queue <- "https://www.example.net/foo"
	queue <- "https://www.example.net/bar"

	close(queue)
	wg.Wait()
}

func fetchURL(queue chan string, i int) {
	for {
		url, more := <-queue
		if more {
			// URL取得
			fmt.Println("fetching", i, url)
			// ...
		} else {
			fmt.Println("worker exit", i)
			wg.Done()
			return
		}
	}
}

func fetchURL02(queue chan string, done chan bool) {
	for url := range queue {
		// URL取得処理
		fmt.Println("URL:", url)
	}
	fmt.Println("worker exit")
	done <- true
}
