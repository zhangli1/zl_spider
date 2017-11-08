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

func (self *Spider) Run() string {
	var content string
    request := NewRequest(self.Cfg)
    content = request.Run()
	parse := NewParse(content)
	return parse.Run()
}


