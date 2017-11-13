package controller

import (
	"time"
	"net/http"
	"bytes"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
    "crypto/tls"
)


type Request struct {
    UserConfig UserConfig
}


func NewRequest(userConfig UserConfig) *Request {
    request := &Request{UserConfig : userConfig}
    return request
}

func (self *Request) Run() *goquery.Document {
    return self.http_request()
}

func (self *Request) http_request() *goquery.Document {
	tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DisableCompression: true,
    }
    timeout_s := self.UserConfig.TimeOut
	c := &http.Client{
	    Timeout: time.Duration(timeout_s) * time.Second,
		Transport: tr,
	}

    var resp *http.Response

    if len(self.UserConfig.Param) < 1 {
        resp, _ = c.Get(self.UserConfig.Url)
    } else {
		b, _ := json.Marshal(self.UserConfig.Param)
		body := bytes.NewBuffer([]byte(b))
        resp, _ = c.Post(self.UserConfig.Url, "", body)
    }
    res, _ := goquery.NewDocumentFromResponse(resp)
    return res
}
