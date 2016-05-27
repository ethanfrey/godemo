package main

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/redis.v3"
	//    "gopkg.in/redsync.v1"  // TODO: use
)

var client *redis.Client

type Post struct {
	Data string `form:"data" json:"data" binding:"required"`
}

var prefix, list_id string

func init() {
	prefix = "goredis:"
	list_id = prefix + "_list"
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println(pong)
	}
}

func index(c *gin.Context) {
	c.String(200, "Check this out!")
}

func getItem(c *gin.Context) {
	item := c.Param("item_id")
	value, err := client.Get(prefix + item).Result()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	} else {
		c.JSON(200, gin.H{"data": value})
	}
}

func setItem(c *gin.Context) {
	item := c.Param("item_id")
	var input Post
	err := c.Bind(&input)
	if err == nil {
		value := input.Data
		key := prefix + item
		err = client.Set(key, value, 0).Err()
		if err == nil {
			err = client.HSet(list_id, key, "1").Err()
		}
	}

	if err != nil {
		c.JSON(400, gin.H{"error": err})
	} else {
		c.JSON(200, gin.H{"success": true})
	}
}

func listItems(c *gin.Context) {
	data, err := client.HGetAllMap(list_id).Result()
	if err != nil {
		c.JSON(400, gin.H{"error": err})
		return
	}

	value := make([]string, 0, 100)
	for k, _ := range data {
		idx := len(value)
		value = value[0 : idx+1]
		value[idx] = strings.TrimPrefix(k, prefix)
	}
	c.JSON(200, gin.H{"data": value})
}

func main() {
	r := gin.Default()

	r.GET("/", index)
	r.GET("/items/:item_id", getItem)
	r.POST("/items/:item_id", setItem)
	r.PUT("/items/:item_id", setItem)
	r.GET("/items", listItems)

	r.Run(":7890") // listen and server on 0.0.0.0:8080
}
