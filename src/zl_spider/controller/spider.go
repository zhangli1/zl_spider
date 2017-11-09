/*
 * zl_spider主入口文件
 */


package controller

import (
	"zl_spider/config"
)

type Spider struct {
    ExeDir string
    Cfg    config.Config
}

func NewSpider(exeDir string, cfg config.Config) *Spider {
    spider := &Spider{ExeDir : exeDir, Cfg : cfg}
    return spider
}

func (self *Spider) Run() interface{} {
    //先从用户规则中读取需要的信息
    var userConfig UserConfig
    rule := NewRule()
    userConfig = rule.Run()
    request := NewRequest(userConfig)
    content := request.Run()
    return content
	//parse := NewParse(content)
	//return parse.Run()
}


