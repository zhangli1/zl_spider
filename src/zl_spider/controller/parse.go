package controller

import ()

type Parse struct {

}

func NewParse() *Parse {
    parse := &Parse{}
    return parse
}

func (self *Parse) Run() interface{} {
    return "haha"
}
