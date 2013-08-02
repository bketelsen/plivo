package plivo

import (
"fmt"
"log"
"io/ioutil"

)

// Struct representing Plivo's SMS parameters.
// Source and Destination are Phone numbers which
// must be prefixed with country code, but no + sign.
// e.g. 18135551232
// Multiple numbers can be used with a delimiter "<"
// e.g. 18135551234<18139661234
// TODO: Destination should be a slice of numbers, not a single number
type SmsMessage struct {
	Source 	string
	Destination	string
	Message string
}

type SmsResponse struct {
	StatusCode	int
	Message 	string
	MessageID	string
}

//TODO: Fix this crap
//TODO: parse response
//TODO: properly json encode 'data' from SmsMessage
func (plivo *Plivo) SendSMS(msg *SmsMessage) error {
	data := `{"src":"` + msg.Source  +`","dst":"` + msg.Destination +  `","text":"`+msg.Message+`"}`
	plivoUrl := BASEURL + "/Account/" + plivo.AuthId + "/Message/"
	log.Println(plivoUrl)
	log.Println(msg)
	resp, err := plivo.post(data,plivoUrl)
	defer resp.Body.Close()
	if err != nil {
		fmt.Println("ERROR:", err)
		return err
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return  err
	}
	// todo parse response and return it

	fmt.Println("HTTP STATUS:"  , resp.StatusCode )
	fmt.Println("Response: ", string(responseBody))
	return err
	
}