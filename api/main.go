package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
)

type UrlRequestBody struct {
    URL string `json:"url"`
}

type Article struct {
    Article string `json:"article"`
}

func ellorM8(ctx *gin.Context) {
    ctx.Header("Access-Control-Allow-Origin", "*")

    ctx.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
}

func getPodcastNotes(ctx *gin.Context) {
    var requestBody UrlRequestBody


    if err := ctx.ShouldBindJSON(&requestBody); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c := colly.NewCollector(
        colly.AllowedDomains("www.podcastnotes.org", "podcastnotes.org"),
    )

    c.OnHTML("article", func(e *colly.HTMLElement) {
        articleHTML, err := e.DOM.Html()
        fmt.Println(articleHTML)
        if err != nil {
            ctx.JSON(http.StatusInternalServerError, gin.H{
                "message": "error",
            })
        }
        ctx.JSON(http.StatusOK, gin.H{
            "article": articleHTML,
        })
    })

    c.Visit(requestBody.URL)
}

func main() {
    const PORT = ":9090"
    r := gin.Default()
    r.GET("/", ellorM8)
    r.GET("/podcastnotes", getPodcastNotes)
    r.Run(PORT)
}
