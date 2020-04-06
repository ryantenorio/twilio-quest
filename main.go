package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const smsURL string = "https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json"

// SMSMessage is a representation of fields required to submit a SMS message request
type SMSMessage struct {
	to                  string
	from                string
	messagingServiceSid string
	body                string
	mediaURL            string
}

// SMSMessageResponse is what is returned after making a POST request for an SMSMessage
type SMSMessageResponse struct {
	AccountSID          string            `json:"account_sid"`
	APIVersion          string            `json:"api_version"`
	Body                string            `json:"body"`
	DateCreated         string            `json:"date_created"`
	DateSent            string            `json:"date_sent"`
	DateUpdated         string            `json:"date_updated"`
	Direction           string            `json:"direction"`
	ErrorCode           string            `json:"error_code"`
	ErrorMessage        string            `json:"error_message"`
	From                string            `json:"from"`
	MessagingServiceSID string            `json:"messaging_service_sid"`
	NumMedia            string            `json:"num_media"`
	NumSegments         string            `json:"num_segments"`
	Price               string            `json:"price"`
	PriceUnit           string            `json:"price_unit"`
	SID                 string            `json:"sid"`
	Status              string            `json:"status"`
	SubresourceUris     map[string]string `json:"subresource_uris"`
	To                  string            `json:"to"`
	URI                 string            `json:"uri,"`
}

func (s SMSMessage) isValid() bool {
	// TODO
	return true
}

type twilioCredentials struct {
	accountSID string
	authToken  string
}

type twilio struct {
	client *http.Client
	twilioCredentials
}

func twilioClient() *twilio {
	return &twilio{
		client: http.DefaultClient,
		twilioCredentials: twilioCredentials{
			accountSID: os.Getenv("TWILIO_ACCOUNT_SID"),
			authToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
		},
	}
}

func (t *twilio) Send(message SMSMessage) (*SMSMessageResponse, error) {
	smsMessageResponse, err := t.send(message)
	if err != nil {
		return nil, err
	}
	return smsMessageResponse, nil
}

func (t *twilio) send(message SMSMessage) (*SMSMessageResponse, error) {
	if !message.isValid() {
		return nil, fmt.Errorf("Message not valid: %s", message)
	}
	values := url.Values{}
	values.Set("To", message.to)
	values.Set("From", message.from)
	values.Set("Body", message.body)
	response, err := t.post(fmt.Sprintf(smsURL, t.accountSID), values)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	smsResponse := &SMSMessageResponse{}
	err = json.Unmarshal(responseBody, smsResponse)
	if err != nil {
		return nil, err
	}
	return smsResponse, nil
}

func (t *twilio) get(url string) (*http.Response, error) {
	request, _ := http.NewRequest("GET", url, nil)
	return t.do(request)
}

func (t *twilio) post(url string, values url.Values) (*http.Response, error) {
	request, _ := http.NewRequest("POST", url, strings.NewReader(values.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return t.do(request)
}

func (t *twilio) do(request *http.Request) (*http.Response, error) {
	request.SetBasicAuth(t.accountSID, t.authToken)
	return t.client.Do(request)
}

// Say is a representation of the TWiML say instruction
type Say string

// Response is a representation of TwiML instructions
type Response struct {
	Says []Say `xml:"say"`
}

func hello(c *gin.Context) {
	// return back a say tag
	response := Response{
		Says: []Say{"Hello, World!", "Welcome to my Twilio App!"},
	}
	c.XML(200, response)
}

func sendSMS(c *gin.Context) {
	client := twilioClient()
	message := SMSMessage{
		to:   "+14805448206",
		from: "+17155022669",
		body: "Test SMS",
	}
	response, err := client.Send(message)
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
