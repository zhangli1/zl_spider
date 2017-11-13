package controller

import (
)

type Rule struct {
}

//需要用户操心的返回结构
type UserConfig struct {
    Url string
    TimeOut int64
    Param map[string] interface{}
}



func NewRule() *Rule {
    rule := &Rule{}
    return rule
}


func (self *Rule) Run() UserConfig {
    var userConfig UserConfig
    userConfig.Url = "https://segmentfault.com/q/1010000010136765"
    //userConfig.Url = "https://www.zhihu.com/question/57171912/answer/256767862"
    //userConfig.Url = "http://test.hqs.haoqiao.cn:8080/log/log_list?server=test&supplier_id=8&date=20171113&hour=-1&action="
    userConfig.TimeOut = 10

    param := make(map[string] interface{})
    param["action"] = "price"
    param["type"]   = 1
    userConfig.Param = param

    return userConfig
}
