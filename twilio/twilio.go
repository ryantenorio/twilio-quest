package twilio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type twilioCredentials struct {
	accountSID string
	authToken  string
}

// Twilio holds a http client and credentials for use with the Twilio API
type Twilio struct {
	client *http.Client
	twilioCredentials
}

// Client returns a new twilio client
func Client() *Twilio {
	return &Twilio{
		client: http.DefaultClient,
		twilioCredentials: twilioCredentials{
			accountSID: os.Getenv("TWILIO_ACCOUNT_SID"),
			authToken:  os.Getenv("TWILIO_AUTH_TOKEN"),
		},
	}
}

// SendSMS sends the given message and returns the response
func (t *Twilio) SendSMS(message SMSMessage) (*SMSMessageResponse, error) {
	smsMessageResponse, err := t.send(message)
	if err != nil {
		return nil, err
	}
	return smsMessageResponse, nil
}

func (t *Twilio) send(message SMSMessage) (*SMSMessageResponse, error) {
	if !message.isValid() {
		return nil, fmt.Errorf("Message not valid: %s", message)
	}
	values := url.Values{}
	values.Set("To", message.To)
	values.Set("From", message.From)
	values.Set("Body", message.Body)
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

func (t *Twilio) get(url string) (*http.Response, error) {
	request, _ := http.NewRequest("GET", url, nil)
	return t.do(request)
}

func (t *Twilio) post(url string, values url.Values) (*http.Response, error) {
	request, _ := http.NewRequest("POST", url, strings.NewReader(values.Encode()))
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	return t.do(request)
}

func (t *Twilio) do(request *http.Request) (*http.Response, error) {
	request.SetBasicAuth(t.accountSID, t.authToken)
	return t.client.Do(request)
}
