package rdl

import (
	"io/ioutil"
	"net/http"
)

type Requester struct {
	Location string
	Options map[string]string
	RequestString string
	Data string
	Errors []error
}

func NewRequester(location string) (nr Requester) {
	nr.Location = location
	nr.Options = make(map[string]string)
	return
}

func (r *Requester) AddOption(key string, value string) {
	r.Options[key] = value
}

func (r *Requester) BuildRequestString() {
	r.RequestString = r.Location
}

func (r *Requester) SendRequest() (success bool) {
	resp, http_err := http.Get(r.RequestString)
	if http_err != nil {
		r.Errors = append(r.Errors, http_err)
		return
	}
	defer resp.Body.Close()

	data, read_err := ioutil.ReadAll(resp.Body)
	if read_err != nil {
		r.Errors = append(r.Errors, read_err)
		return
	}

	r.Data = string(data)

	success = true
	return
}

