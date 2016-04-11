package main

import "github.com/gin-gonic/gin"

var data map[string]string

type Post struct {
    Data     string `form:"data" json:"data" binding:"required"`
}


func init() {
    data = map[string]string{
        "hi": "dude",
        "hello": "world",
        "what": "when",
    }
}

func index(c *gin.Context) {
    c.String(200, "Check this out!")
}

func getItem(c *gin.Context) {
    item := c.Param("item_id")
    value := data[item]
    c.String(200, value)
}

func setItem(c *gin.Context) {
    item := c.Param("item_id")
    var input Post
    err := c.Bind(&input)
    if err == nil {
        value := input.Data
        data[item] = value
        c.JSON(200, gin.H{"success": true})
    } else {
        c.JSON(400, gin.H{"error": err})
    }
}

func listItems(c *gin.Context) {
    c.JSON(200, data)
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
