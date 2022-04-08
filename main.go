// 2022/4/7
// zhangzhong
// a very simple kv server

package main

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// client request: GET /get  get all the keys
// client request: GET /get/key
// client request: POST /put {"key": "value", value: "value"}
// client request: POST /append {"key": "value", value: "value"}

// we have a in-memory map and a lock to protect it
type kvserver struct {
	m    sync.RWMutex
	data map[string]string
}

// json must export!! the field name must be Captial Word!!!
type kvpair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func main() {
	// init kv server
	var kv kvserver
	kv.data = make(map[string]string)
	kv.data["author"] = "zhangzhong"

	router := gin.Default()
	router.GET("/get", func(ctx *gin.Context) {
		kv.m.RLock()
		defer kv.m.RUnlock()

		ctx.IndentedJSON(http.StatusOK, kv.data)
	})

	router.GET("/get/:key", func(ctx *gin.Context) {
		kv.m.RLock()
		defer kv.m.RUnlock()

		key := ctx.Param("key")
		value, ok := kv.data[key]
		if !ok {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"key": "not found"})
			return
		}
		ctx.IndentedJSON(http.StatusOK, gin.H{key: value})
	})

	router.POST("/put", func(ctx *gin.Context) {
		kv.m.Lock()
		defer kv.m.Unlock()

		var pair kvpair
		ctx.BindJSON(&pair)
		kv.data[pair.Key] = pair.Value
		ctx.IndentedJSON(http.StatusCreated, gin.H{pair.Key: pair.Value})
	})

	router.POST("/append", func(ctx *gin.Context) {
		kv.m.Lock()
		defer kv.m.Unlock()

		var pair kvpair
		ctx.BindJSON(&pair)
		kv.data[pair.Key] = kv.data[pair.Key] + pair.Value
		ctx.IndentedJSON(http.StatusCreated, gin.H{pair.Key: kv.data[pair.Key]})
	})

	router.Run("0.0.0.0:8080")
}
