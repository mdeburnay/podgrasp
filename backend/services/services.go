package services

import (
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

type UrlRequestBody struct {
    URL string `json:"url"`
}

type PodcastArticle struct {
    Date string `json:"date"`
    Title string `json:"title"`
    Headers map[int]string `json:"headers"`
    Sections map[int]interface{} `json:"sections"`
}

func EllorM8(ctx *gin.Context) {
    ctx.Header("Access-Control-Allow-Origin", "*")

    ctx.JSON(http.StatusOK, gin.H{
      "message": "pong",
    })
}

func SendEmail(ctx *gin.Context) {

    err := godotenv.Load("../.env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

    requestBody := UrlRequestBody{}

    if err := ctx.ShouldBindJSON(&requestBody); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    fmt.Println("Sending...")

    PodcastHTML := GetPodcastNotes(requestBody.URL);

    if err != nil {
        log.Fatal("Failure getting podcast notes: ", err)
    }

    FormattedArticle := FormatEmail(PodcastHTML)

    if err != nil {
        log.Fatal("Failure formatting email: ", err)
    }

    fmt.Println(FormattedArticle.Headers)

    // sendgridKey := os.Getenv("SENDGRID_API_KEY")

    // from := mail.NewEmail("Podgrasp", "YOUR_SENDGRID_VERIFIED_SENDER_EMAIL_ADDRESS")
    // subject := "New Podcast Notes From ${Insert Podcast Name Here}"
    // to := mail.NewEmail("Example User", "USER_EMAIL_ADDRESS")
    // plainTextContent := "Boop beep bap"
    // htmlContent := PodcastHTML
    // message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
    // client := sendgrid.NewSendClient(sendgridKey)
    // response, err := client.Send(message)
    // if err != nil {
    //  log.Println(err)
    // } else {
    //  fmt.Println(response.StatusCode)
    //  fmt.Println(response.Body)
    //  fmt.Println(response.Headers)
    // }
}

func GetPodcastNotes(podcastURL string) (PodcastArticle) {

    headers := make(map[int]string)

    sections := make(map[int]interface{})

    c := colly.NewCollector(
        colly.AllowedDomains("www.podcastnotes.org", "podcastnotes.org"),
    )

    c.OnHTML("article", func(e *colly.HTMLElement) {

        e.DOM.Find("header")

        e.ForEach("h4", func(i int, e *colly.HTMLElement) {
            headers[i] = e.Text
        })

        e.ForEach("ul", func(i int, elem *colly.HTMLElement) {
            section := make(map[int]string)
            elem.ForEach("li", func(i int, e *colly.HTMLElement) {
                section[i] = e.Text
            })
            sections[i] = section
        })
    })

    c.Visit(podcastURL)

    finishedPodcast := PodcastArticle{
        Date: "2020-01-01",
        Title: "Test",
        Headers: headers,
        Sections: sections,
    }

    return finishedPodcast
}

func FormatEmail(article PodcastArticle) (PodcastArticle) {
    articleHeaders := article.Headers
    articleSections := article.Sections

    // Sort Article Headers
    headerKeys := make([]int, 0, len(articleHeaders))

    for k := range articleHeaders {
        headerKeys = append(headerKeys, k)
    }

    sort.Ints(headerKeys)

    // Sort Article Sections
    sectionKeys := make([]int, 0, len(articleSections))

    for k := range articleSections {
        sectionKeys = append(sectionKeys, k)
    }

    sort.Ints(sectionKeys)

    formattedArticle := PodcastArticle{
        Date: article.Date,
        Title: article.Title,
        Headers: articleHeaders,
        Sections: articleSections,
    }

    return formattedArticle
}
