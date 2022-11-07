package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)

type UrlRequestBody struct {
    URL string `json:"url"`
}

type Article struct {
    Article string `json:"article"`
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

    var requestBody UrlRequestBody

    if err := ctx.ShouldBindJSON(&requestBody); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    
    var PodcastHTML = GetPodcastNotes(requestBody.URL);

    var FormattedEmail = FormatEmail(PodcastHTML);

    fmt.Println(FormattedEmail)
    
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

func GetPodcastNotes(podcastURL string) (string) {
	fmt.Println("Getting podcast notes for: " + podcastURL)
    var article string
    c := colly.NewCollector(
        colly.AllowedDomains("www.podcastnotes.org", "podcastnotes.org"),
    )

    c.OnHTML("article", func(e *colly.HTMLElement) {
        articleHTML, err := e.DOM.Html()
        if err != nil {
            log.Fatal(err)
        }
        article = articleHTML
    })

    c.Visit(podcastURL)

    return article
}

func FormatEmail(article string) (string) {
    return article
}
