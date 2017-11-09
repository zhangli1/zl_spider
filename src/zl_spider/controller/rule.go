package controller

import (
)

type Rule struct {
}

//需要用户操心的返回结构
type UserConfig struct {
    Url string
    TimeOut int
    Param map[string] interface{}
}



func NewRule() *Rule {
    rule := &Rule{}
    return rule
}


func (self *Rule) Run() UserConfig {
    var userConfig UserConfig
    userConfig.Url = "https://www.zhihu.com/question/57171912/answer/256767862"
    userConfig.TimeOut = 10

    param := make(map[string] interface{})
    param["action"] = "price"
    param["type"]   = 1
    userConfig.Param = param

    return userConfig
}
