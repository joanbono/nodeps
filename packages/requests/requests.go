package requests

import (
	"bytes"
	"crypto/tls"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

/*
curl -k -x "http://127.0.0.1:9090" -X POST "https://api.npms.io/v2/package/mget" \
	    -H "Accept: application/json" \
	    -H "Content-Type: application/json" \
	    -d '["@aws-amplify/ui", "@xstate/react", "xstate"]'
*/

// Add proxy for debugging
var proxyUrl, err = url.Parse("http://127.0.0.1:9090")

var tr = &http.Transport{
	TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	Proxy:           http.ProxyURL(proxyUrl),
}

var client = &http.Client{Timeout: time.Second * 10, Transport: tr}

func MakeRequest(packageList string) (body []byte) {
	url := `https://api.npms.io/v2/package/mget`

	var jsonStr = []byte(packageList)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-Applicagion", "nodeps")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", `Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:62.0) Gecko/20100101 Firefox/62.0`)
	req.Header.Set("Content-Type", "application/json")

	//client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
