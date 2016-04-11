package main

import (
    "html/template"
    "strings"
    "github.com/gin-gonic/gin"
)

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
//        "data": map[string]string{
        "data": gin.H{
            "foo": "bar",
            "more": "stuff",
            "for": "you",
        },
    })
}

func privateData(c *gin.Context) {
    c.String(200, "private data")
}

// Modified code from gin.LoadHTMLGlob
// see https://github.com/gin-gonic/gin/issues/449
func loadHTMLGlob(engine *gin.Engine, pattern string, funcMap template.FuncMap) {
    templ := template.Must(template.New("").Funcs(funcMap).ParseGlob(pattern))
    engine.SetHTMLTemplate(templ)
}

func main() {
    r := gin.Default()
    setUpStatic(r)

    templateFilters := template.FuncMap{
        "upper": strings.ToUpper, // A custom template function
    }

    loadHTMLGlob(r, "templates/static/*.html", templateFilters)
    r.GET("/", index)
    r.GET("/info", info)

    private := r.Group("/private")
    private.GET("data", privateData)

    // r.Run(":7890")
    certFile := "./certs/certificate.pem"
    keyFile := "./certs/localhost_key.pem"
    r.RunTLS(":7891", certFile, keyFile)
}

/*
Making certs...
https://certsimple.com/blog/localhost-ssl-fix

openssl pkcs12 -in Certificates.p12 -out Certificates.pem -nodes
openssl x509 -inform der -in localhost_cert.cer -out certificate.pem

*/
