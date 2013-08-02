package plivo

import (

)

const BASEURL = "https://api.plivo.com/v1"

type Plivo struct {
	AuthId	string
	AuthToken string
}


// Create a new Twilio struct.
func NewPlivoClient(authId, authToken string) *Plivo {
	return &Plivo{authId, authToken}
}
