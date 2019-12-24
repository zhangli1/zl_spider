/*
 * zl_spider主入口文件
 */

package controller

import (
	"fmt"
	"zl_spider/config"
)

type Spider struct {
	ExeDir string
	Cfg    config.Config
}

func NewSpider(exeDir string, cfg config.Config) *Spider {
	spider := &Spider{ExeDir: exeDir, Cfg: cfg}
	return spider
}

func (self *Spider) Run() {
	//读取配置和规则
	rule := NewRule()
	user_config_info_list := rule.Run(self.Cfg)

	for _, v := range user_config_info_list {
		if v.Switch == false {
			continue
		}
		route := NewRoute(v, self.Cfg)
		fmt.Println(route.Run())
	}
}

/*func (self *Spider) GetInfo(info config.UserConfigInfo) string {
	//    time.Sleep(10 * 1000 * time.Millisecond)
	//进行请求
	//var content *goquery.Document
	request := NewRequest(info)
	content := request.Run()

	//进行解析
	parse := NewParse(content, info, self.Cfg)
	jsons, _ := json.Marshal(parse.Run()) //转换成JSON返回的是byte[]
	return string(jsons)

	return ""
}*/
