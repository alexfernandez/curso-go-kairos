package main

import (
	"memcached"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go memcached.Start()
	wg.Wait()

}
