/*
 * zl_spider主入口文件
 */

package controller

import (
	"os"
	"sync"
	"zl_spider/config"

	l4g "code.google.com/p/log4go"
)

type Spider struct {
	ExeDir    string
	Cfg       config.Config
	LoggerMap map[string]*l4g.Logger
	Sigs      chan os.Signal
	Running   bool
}

func NewSpider(exeDir string, cfg config.Config, logger_map map[string]*l4g.Logger, sigs chan os.Signal) *Spider {
	spider := &Spider{ExeDir: exeDir, Cfg: cfg, LoggerMap: logger_map, Sigs: sigs}
	return spider
}

func (self *Spider) Run() {
	self.Running = false
	//读取配置和规则
	rule := NewRule()
	user_config_info_list := rule.Run(self.Cfg)

	var wg sync.WaitGroup

	//runTask := make([]config.UserConfigInfo, 0)
	for _, v := range user_config_info_list {
		if v.Switch == false {
			continue
		}
		wg.Add(1)
		//runTask = append(runTask, v)
		go NewRoute(v, self.Cfg, self.LoggerMap, self.Sigs, wg).Run()
	}

	/*for _, vv := range runTask {
		go NewRoute(vv, self.Cfg, self.l4gLogger, self.Sigs, wg).Run()
	}*/
	wg.Wait()

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
