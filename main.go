package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ryantenorio/twilio-quest/twilio"
)

func hello(c *gin.Context) {
	// return back a say tag
	response := twilio.Response{
		Says: []twilio.Say{"Hello, World!", "Welcome to my Twilio App!"},
	}
	c.XML(200, response)
}

func sendSMS(c *gin.Context) {
	client := twilio.Client()
	message := twilio.SMSMessage{
		To:   os.Getenv("TO_TEST"),
		From: os.Getenv("TWILIO_NUMBER"),
		Body: "Test SMS",
	}
	response, err := client.SendSMS(message)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}
	c.JSON(200, response)
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	router.POST("/hello", hello)
	router.POST("/sendTestSMS", sendSMS)

	router.Run(":5000")
}
