package controller

import ()

type Out struct {

}

func NewOut() *Out {
    out := &Out{}
    return out
}

func (self *Out) Run() interface{} {
    return "out"
}
