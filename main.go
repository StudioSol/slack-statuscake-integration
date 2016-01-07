package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
	"os"
)

type SlackAttachments struct {
	Color    string `json:"color"`
	Fallback string `json:"fallback"`
	Title    string `json:"title"`
	Text     string `json:"text"`
}

type SlackParams struct {
	Attachments []SlackAttachments `json:"attachments"`
	Username    string             `json:"username"`
	Mrkdwn      bool               `json:"mrkdwn"`
	IconEmoji   string             `json:"icon_emoji"`
}

func color(s string) string {
	if s == "Up" {
		return "#00fd55"
	}
	return "#cc3300"
}

func text(url string) string {
	return fmt.Sprintf("<%s|Link> • <https://www.statuscake.com/App/YourStatus.php|View details>", url)
}

func fallback(name, status string) string {
	return fmt.Sprintf("%s is %s - https://www.statuscake.com/App/YourStatus.php", name, status)
}

func title(name, status string) string {
	return fmt.Sprintf("%s is %s", name, status)
}

func handler(c *gin.Context) {
	URL := c.PostForm("URL")
	name := c.PostForm("Name")
	status := c.PostForm("Status")
	slack := SlackParams{
		Attachments: []SlackAttachments{
			SlackAttachments{
				Color:    color(status),
				Fallback: fallback(name, status),
				Title:    title(name, status),
				Text:     text(URL),
			},
		},
		Username:  "StatusCake Bot",
		Mrkdwn:    true,
		IconEmoji: ":cake:",
	}
	url_slack := fmt.Sprintf("https://hooks.slack.com/%s", c.Param("url"))
	b, err := json.Marshal(slack)
	if err != nil {
		c.String(400, "fail")
		return
	}
	http.PostForm(url_slack, url.Values{"payload": {string(b)}})
	c.String(200, "ok")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := gin.Default()
	router.POST("/*url", handler)
	router.Run(":" + port)
}
