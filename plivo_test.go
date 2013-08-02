package plivo

import (
"testing"
)

func TestSendSMS(t *testing.T) {
	plivo := NewPlivoClient("YOURAPIID","YOURAUTHCODE")
	msg := &SmsMessage{"18135551212","18135551212","Testing From Plivo"}
	plivo.SendSMS(msg)
}