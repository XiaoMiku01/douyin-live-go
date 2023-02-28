package main

import (
	"sync"
)

func main() {
	r, err := NewRoom("https://live.douyin.com/7003418886")
	if err != nil {
		panic(err)
	}
	r.Connect()
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
