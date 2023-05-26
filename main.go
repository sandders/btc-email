package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"strconv"
	"time"
)

var sender = ""
var recipient = ""
var smtpHost = ""
var smtpPort int
var smtpUser = ""
var smtpPassword = ""

func main() {
	var err error
	sender = os.Getenv("EMAIL_SENDER")
	smtpHost = os.Getenv("SMTP_HOST")
	smtpPort, err = strconv.Atoi(os.Getenv("SMTP_PORT"))
	smtpUser = os.Getenv("SMTP_USER")
	smtpPassword = os.Getenv("SMTP_PASSWORD")

	// Check if any of the environment variables are missing
	if sender == "" || smtpHost == "" || err != nil || smtpUser == "" || smtpPassword == "" {
		fmt.Printf("One or more required environment variables are missing or set wrong value type. %s", err)
		return
	}

	client = &http.Client{Timeout: 10 * time.Second}
	router := gin.Default()

	router.GET("/rate", getRate)
	router.POST("/subscribe", subscribe)
	router.POST("/sendEmails", sendEmailsHandler)

	router.Run("0.0.0.0:8080")
}
