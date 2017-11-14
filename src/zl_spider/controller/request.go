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
	var req *http.Request

	b, _ := json.Marshal(self.UserConfig.Param)
	body := bytes.NewBuffer([]byte(b))
    if len(self.UserConfig.Param) < 1 {
        //resp, _ = c.Get(self.UserConfig.Url)
		req, _ = http.NewRequest("GET", self.UserConfig.Url, nil)
    } else {
        //resp, _ = c.Post(self.UserConfig.Url, "application/x-www-form-urlencoded", body)

		req, _ = http.NewRequest("POST", self.UserConfig.Url, body)
    }
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_4) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/61.0.3163.100 Safari/537.36")
	resp, _ = c.Do(req)

    res, _ := goquery.NewDocumentFromResponse(resp)
    return res
}
