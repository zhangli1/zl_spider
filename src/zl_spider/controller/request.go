package controller

import ()


type Request struct {
}


func NewRequest() *Request {
    request := &Request{}
    return request
}

func (self *Request) Run() interface{} {
    return "request"
}
