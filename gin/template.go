package main

import "github.com/gin-gonic/gin"

func setUpStatic(r *gin.Engine) {
    r.Static("/js", "./js")
    r.Static("/css", "./css")
    r.Static("/img", "./img")
    r.StaticFile("/favicon.ico", "./img/favicon.ico")
}

func index(c *gin.Context) {
    c.HTML(200, "index.html", gin.H{
        "title": "My Awesome Index",
        "current": "index",
        })
}

func info(c *gin.Context) {
    c.HTML(200, "info.html", gin.H{
        "title": "Some special info",
        "current": "info",
        })
}

func privateData(c *gin.Context) {
    c.String(200, "private data")
}

func main() {
    r := gin.Default()
    setUpStatic(r)

    r.LoadHTMLGlob("templates/*.html")
    r.GET("/", index)
    r.GET("/info", info)

    private := r.Group("/private")
    private.GET("data", privateData)

    r.Run(":7890")
//    r.RunTLS(":7891")
}
