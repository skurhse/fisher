package lrucache

import (
	"encoding/json"
	"fmt"
	"strings"
)

func ExampleLRUCacheOne() {

	actions := []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"}
	values := [][]int{{2}, {1, 1}, {2, 2}, {1}, {3, 3}, {2}, {4, 4}, {1}, {3}, {4}}

	ExemplifyLRUCache(actions, values)
	// Output:
	//[null, null, null, 1, null, -1, null, -1, 3, 4]
}

func ExemplifyLRUCache(actions []string, values [][]int) {
	var lruCache *LRUCache
	output := make([]any, len(actions))
	for i, action := range actions {
		switch action {
		case "LRUCache":
			lruCache = New(values[i][0])
		case "put":
			lruCache.Put(values[i][0], values[i][1])
		case "get":
			output[i] = lruCache.Get(values[i][0])
		}
	}

	json, err := json.Marshal(output)
	if err != nil {
		panic(err)
	}

	fmt.Println(strings.ReplaceAll(string(json), ",", ", "))
}
