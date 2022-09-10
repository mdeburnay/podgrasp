package main

import (
	"fmt"
	"net/http"

	"github.com/gocolly/colly"
)

func setupRoutes() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Simple Server")
    })
}

func main() {

    setupRoutes()
    http.ListenAndServe(":8080", nil)
    c := colly.NewCollector()

    c.OnHTML("html", func(e *colly.HTMLElement) {
        fmt.Println(e.Text)
    })

    c.Visit("https://podcastnotes.org/huberman-lab/episode-84-sleep-toolkit-tools-for-optimizing-sleep-sleep-wake-timing-huberman-lab/")
}
