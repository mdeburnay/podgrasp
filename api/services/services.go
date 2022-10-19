package services

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
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
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	sendgridKey := os.Getenv("SENDGRID_API_KEY")

	from := mail.NewEmail("Example User", "de_burnay@hotmail.co.uk")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Example User", "deburnayb@gmail.com")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(sendgridKey)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

func GetPodcastNotes(ctx *gin.Context) {
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
