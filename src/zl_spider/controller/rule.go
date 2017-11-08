package controller

import ()

type Rule struct {
}

func NewRule() *Rule {
    rule := &Rule{}
    return rule
}


func (self *Rule) Run() interface{} {
    return "rule"
}
