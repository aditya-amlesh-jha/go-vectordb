package algorithms

import (
	"github.com/iancoleman/orderedmap"
)

type LRUCache struct {
	memory *orderedmap.OrderedMap
}

func LRUCacheObject() *LRUCache {
	return &LRUCache{
		memory: orderedmap.New(),
	}
}

type LRUFunctions interface {
	Get(key string) string
	Put(key string, value string) (string, string)
	Delete(key string) string
	freeMemory(time int)
}

func (cache *LRUCache) Get(key string) string {
	if len(key) == 0 {

	}
}

func (cache *LRUCache) Put(key string) (string, string) {
	if len(key) == 0 {

	}
}

func (cache *LRUCache) Delete(key string) (string, string) {
	if len(key) == 0 {

	}
}

func (cache *LRUCache) freeMemory(key string) (string, string) {
	if len(key) == 0 {

	}
}
