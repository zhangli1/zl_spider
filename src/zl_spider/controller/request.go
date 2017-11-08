package controller

import (
    "zl_spider/config"
	"net/http"
	"io/ioutil"
)


type Request struct {
    Cfg config.Config
}


func NewRequest(cfg config.Config) *Request {
    request := &Request{Cfg : cfg}
    return request
}

func (self *Request) Run() string {
    return self.get()
}

func (self *Request) get() string {
	resp, err := http.Get(self.Cfg.Base.Url)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	return string(body)
}
