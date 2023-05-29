package lrucache

import (
	"errors"
	"fmt"
	"sync"
)

type Cache interface {
	Put(key string, value string) bool
	Get(key string) (string, error)
}

type lruCache struct {
	capacity    int
	keyValueMap map[string]*ListNode
	head        *ListNode
	tail        *ListNode
	mutex       sync.RWMutex
}

func NewLruCache(capacity int) Cache {
	head := &ListNode{}
	tail := &ListNode{}
	head.right = tail
	tail.left = head

	return &lruCache{
		capacity:    capacity,
		keyValueMap: make(map[string]*ListNode),
		head:        head,
		tail:        tail,
	}
}

func (lrc *lruCache) Get(key string) (string, error) {
	lrc.mutex.Lock()
	defer lrc.mutex.Unlock()

	node, exist := lrc.keyValueMap[key]

	if !exist {
		return "", errors.New(fmt.Sprintf("key %s does not exit", key))
	}

	lrc.moveFront(node)

	return node.value, nil
}

func (lrc *lruCache) Put(key string, value string) bool {
	lrc.mutex.Lock()
	defer lrc.mutex.Unlock()

	if len(lrc.keyValueMap) >= lrc.capacity {
		tailKey := lrc.tail.left.key
		DeleteListNode(lrc.tail.left)
		delete(lrc.keyValueMap, tailKey)
	}

	node := &ListNode{key: key, value: value}
	lrc.moveFront(node)

	lrc.keyValueMap[key] = node

	return true
}

func (lrc *lruCache) moveFront(node *ListNode) {
	curFront := lrc.head.right

	lrc.head.right = node
	node.left = lrc.head
	node.right = curFront
	curFront.left = node
}
