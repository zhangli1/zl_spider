/*
 * zl_spider主入口文件
 */


package controller

import (
	"zl_spider/config"
    "github.com/PuerkitoBio/goquery"
//    "time"
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

    //读取配置和规则
    rule := NewRule()
    user_config_info_list := rule.Run(self.Cfg)

    c := make(chan string, len(user_config_info_list))

    ret_info := make(map[int] interface{})
    for _, info := range user_config_info_list {
        go self.GetInfo(info, c)
    }

    for i, _ := range user_config_info_list {
        ret_info[i] = <- c
    }
    return ret_info
}

func (self *Spider) GetInfo(info config.UserConfigInfo, c chan string) {
//    time.Sleep(10 * 1000 * time.Millisecond)
    //进行请求
    var content *goquery.Document
    request := NewRequest(info)
    content = request.Run()

    //进行解析
    parse := NewParse(content)
    c <- parse.Run().(string)
}



