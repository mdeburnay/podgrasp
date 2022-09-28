package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
    const PORT = ":9090"
    r := gin.Default()
    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
          "message": "pong",
        })
      })
      r.Run(PORT)
}


// c := colly.NewCollector()

//     c.OnHTML("article", func(e *colly.HTMLElement) {
//         fmt.Println(e.Text)
//     })

//     c.Visit("https://podcastnotes.org/huberman-lab/episode-84-sleep-toolkit-tools-for-optimizing-sleep-sleep-wake-timing-huberman-lab/")
