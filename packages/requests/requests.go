package requests

import (
	"crypto/tls"
	"log"
	"net/http"
)

// MakeRequest will do the GET request
func MakeRequest(host string) *http.Response {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	response, err := client.Get(host)
	if err != nil {
		log.Fatal(err)
	}

	return response
}
