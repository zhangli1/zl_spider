/*
 * zl_spider主入口文件
 */


package controller

import (
)

type Spider struct {
    ExeDir string
    Cfg    string
}

func NewSpider(exeDir string, cfg string) *Spider {
    spider := &Spider{ExeDir : exeDir, Cfg : cfg}
    return spider
}

func (self *Spider) Run() interface{} {
    request := NewRequest()
    return request.Run()
}
