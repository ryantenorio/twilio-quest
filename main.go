package main

import (
	"github.com/gin-gonic/gin"
)

// Say is a representation of the TWiML say instruction
type Say string

// Response is a representation of TwiML instructions
type Response struct {
	Says []Say `xml:"say"`
}

func hello(c *gin.Context) {
	// return back a say tag
	response := Response{
		Says: []Say{"Hello, World!"},
	}
	c.XML(200, response)
}

func main() {
	// initialize the twilio credentials. We will need this later!
	// TwilioAccountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	// TwilioAuthToken := os.Getenv("TWILIO_AUTH_TOKEN")

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "hello world")
	})
	router.POST("/hello", hello)

	router.Run(":5000")
}
