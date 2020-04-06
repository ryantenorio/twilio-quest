package twilio

const smsURL string = "https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json"

// SMSMessage is a representation of fields required to submit a SMS message request
type SMSMessage struct {
	To                  string
	From                string
	MessagingServiceSid string
	Body                string
	MediaURL            string
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
