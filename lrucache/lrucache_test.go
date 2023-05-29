package lrucache

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	wg           sync.WaitGroup
	testLruCache = NewLruCache(4)
	writeFunc    = func(key string, value string) {
		defer wg.Done()
		testLruCache.Put(key, value)
	}
	readFunc = func(key string) {
		defer wg.Done()
		testLruCache.Get(key)
	}
)

func TestPut(t *testing.T) {
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		str := fmt.Sprintf("%d", i)
		go writeFunc(str, str)
	}

	wg.Wait()

	// 1 will be evicted when inserting 5

	_, err := testLruCache.Get("1")
	assert.NotNil(t, err)

	expVal := "4"
	val, err := testLruCache.Get("4")
	assert.Nil(t, err)
	assert.Equal(t, expVal, val)
}

func TestGet(t *testing.T) {
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		str := fmt.Sprintf("%d", i)
		go writeFunc(str, str)
	}
	wg.Wait()

	wg.Add(1)
	go readFunc("3")
	wg.Add(1)
	go readFunc("3")
	wg.Add(1)
	go readFunc("2")
	wg.Wait()

	testLruCache.Put("6", "6")

	// 2 and 3 will not be evicted when inserting 6

	expVal := "3"
	val, err := testLruCache.Get("3")
	assert.Nil(t, err)
	assert.Equal(t, expVal, val)

	expVal = "2"
	val, err = testLruCache.Get("2")
	assert.Nil(t, err)
	assert.Equal(t, expVal, val)
}
