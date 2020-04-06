package twilio

// Say is a representation of the TWiML say instruction
type Say string

// Response is a representation of TwiML instructions
type Response struct {
	Says []Say `xml:"say"`
}
