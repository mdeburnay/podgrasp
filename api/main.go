package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {


    c := colly.NewCollector()

    c.OnHTML("article", func(e *colly.HTMLElement) {
        fmt.Println(e.Text)
    })

    c.Visit("https://podcastnotes.org/huberman-lab/episode-84-sleep-toolkit-tools-for-optimizing-sleep-sleep-wake-timing-huberman-lab/")
}
