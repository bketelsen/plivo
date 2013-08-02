package plivo

import ( 
	"net/http"
	"strings"
)



func (plivo *Plivo) post(msg, plivoUrl string) (*http.Response, error) {
	req, err := http.NewRequest("POST", plivoUrl, strings.NewReader(msg))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(plivo.AuthId, plivo.AuthToken)
	req.Header.Add("Content-Type", "application/json")

	client := &http.Client{}
	return client.Do(req)
}