package test

import (
	"testing"
	"net/url"
	"net/http"
	"strings"
	"io/ioutil"
	"log"
)

func TestDetailBerat(t *testing.T) {
    // Initialize the request address and request parameters
    uri := "http://localhost:9008/detail_berat"

	formData := url.Values{
        "params[tanggal]": {"2018-08-23"},
    }

    client := &http.Client{}
    
    //Not working, the post data is not a form
    req, err := http.NewRequest("POST", uri, strings.NewReader(formData.Encode()))
    if err != nil {
        log.Fatalln(err)
    }
    
    req.Header.Set("User-Agent", "Golang_Super_Bot/0.1")
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    
    resp, err := client.Do(req)
    if err != nil {
        log.Fatalln(err)
    }
    defer resp.Body.Close()
    
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln(err)
    }
    
    log.Println(string(body))
}