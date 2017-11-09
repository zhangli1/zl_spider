package controller

import (
	"time"
	"net/http"
	"bytes"
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
)


type Request struct {
    UserConfig UserConfig
}


func NewRequest(userConfig UserConfig) *Request {
    request := &Request{UserConfig : userConfig}
    return request
}

func (self *Request) Run() interface{} {
    return self.request()
}

func (self *Request) request() interface{} {
	timeout := self.UserConfig.TimeOut
	c := &http.Client{
	    Timeout: timeout * time.Second,
	}

    if len(self.UserConfig.Param) < 1 {
	    resp, _ := c.Get(self.UserConfig.Url)
    } else {
		b, err := json.Marshal(self.UserConfig.Param)
		if err != nil {
			return err
		}

		body := bytes.NewBuffer([]byte(b))
        c.Post(self.UserConfig.Url, "", &buf)
    }
    return goquery.NewDocumentFromResponse(resp)
}
