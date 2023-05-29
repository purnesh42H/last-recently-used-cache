package main

import (
	"fmt"
	"last-recently-used-cache/lrucache"
	"sync"
)

func main() {
	lruCache := lrucache.NewLruCache(4)

	var wg sync.WaitGroup

	writeFunc := func(key string, value string) {
		defer wg.Done()
		lruCache.Put(key, value)
	}

	readFunc := func(key string) string {
		defer wg.Done()

		value, err := lruCache.Get(key)
		if err != nil {
			fmt.Printf("error from cache %v\n", err)
		} else {
			fmt.Printf("fetched value %s of key %s\n", value, key)
		}

		return value
	}

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		str := fmt.Sprintf("%d", i)
		go writeFunc(str, str)
	}

	wg.Wait()

	// 1 will be evicted from lrucache
	fmt.Printf("Test1\n")

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		str := fmt.Sprintf("%d", i)
		go readFunc(str)
	}

	wg.Wait()

	wg.Add(1)
	go readFunc("3")
	wg.Add(1)
	go readFunc("3")
	wg.Add(1)
	go readFunc("2")

	wg.Wait()

	wg.Add(1)
	go writeFunc("6", "6")

	wg.Wait()

	// 3 or 4 or 5 will be evicted from lrucache
	fmt.Printf("Test2\n")

	for i := 1; i <= 6; i++ {
		wg.Add(1)
		str := fmt.Sprintf("%d", i)
		go readFunc(str)
	}

	wg.Wait()

}
