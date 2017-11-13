/*
 * zl_spider主入口文件
 */


package controller

import (
	"zl_spider/config"
    "github.com/PuerkitoBio/goquery"
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

    //读取配置和规则
    rule := NewRule()
    userConfig = rule.Run()

    //进行请求
    var content *goquery.Document
    request := NewRequest(userConfig)
    content = request.Run()

    //进行解析
	parse := NewParse(content)
	return parse.Run()
}


