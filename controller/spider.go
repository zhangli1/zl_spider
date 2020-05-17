/*
 * zl_spider主入口文件
 */

package controller

import (
	"encoding/json"
	"fmt"
	glib "lib"
	"math/rand"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"
	"sync"
	"time"
	"zl_spider/config"
	"zl_spider/model"

	l4g "code.google.com/p/log4go"
)

type Spider struct {
	ExeDir    string
	Cfg       config.Config
	LoggerMap map[string]*l4g.Logger
	Sigs      chan os.Signal
	Running   bool
	model.FreeIpModel
}

var FilterIPListWriteIndex int

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

	//初始化数据库
	var mysqlconn string
	mysqlconn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s", self.Cfg.Db.UserName, self.Cfg.Db.Passwd, self.Cfg.Db.Host, self.Cfg.Db.Port, self.Cfg.Db.Database, self.Cfg.Db.Charset)
	self.Cfg.MYSQL = glib.NewSQL(mysqlconn, self.LoggerMap["freeip"])
	self.Cfg.MYSQL.Init()

	redis := glib.NewRedis(self.Cfg.Redis.Host, self.Cfg.Redis.Port, self.Cfg.Redis.Passwd, self.Cfg.Redis.Select)

	//runTask := make([]config.UserConfigInfo, 0)
	for _, v := range user_config_info_list {
		if v.Switch == false {
			continue
		}
		wg.Add(1)
		//runTask = append(runTask, v)
		go NewRoute(v, self.Cfg, self.LoggerMap, self.Sigs, wg).Run()
	}

	//代理黑名单写入磁盘
	wg.Add(1)
	go func() {
		for {
			flength := len(model.FilterIPList)
			if len(model.FilterIPList) > 0 {
				blackFileName := "log/blacklist"
				blacklist := glib.ReadFile(blackFileName)
				blacklistArr := strings.Split(blacklist, "\n")
				for i, v := range model.FilterIPList {
					if i < FilterIPListWriteIndex {
						continue
					}
					if !glib.IsExistByKey(v, blacklistArr) {
						glib.WriteFile(blackFileName, fmt.Sprintf("%s\n", v))
						FilterIPListWriteIndex = i
						fmt.Println(fmt.Sprintf("write blacklist: %s", v))
					}
				}
				model.FilterIPList = model.FilterIPList[flength:]
			}
			time.Sleep(time.Second * time.Duration(5))
		}
	}()

	//检查过的代理写入redis
	wg.Add(1)
	go func() {
		for {
			time.Sleep(time.Second * time.Duration(5))
			self.LoggerMap["freeip"].Info("check checkSucProxyList", model.SucData)
			checkSucProxyList := make([]string, 0)

			if len(model.SucData) < 1 {
				continue
			}

			for key, value := range model.SucData {
				//取最后一百个
				lastList := make([]string, 0)
				count := len(value.List)

				max := count
				if count > 100 {
					max = 100
				}
				for i := 0; i < max; i++ {
					index := count - i - 1
					lastList = append(lastList, value.List[index])
				}

				value.List = lastList

				rand.Seed(time.Now().UnixNano()) //随机干扰
				for _, v := range lastList {
					ret := false
					glib.Try(func() {
						reqProxy := v
						ret = self.FreeIpModel.Curl(reqProxy, value.Url, value.CheckStr)
						if ret {
							checkSucProxyList = append(checkSucProxyList, reqProxy)
						}
					}, func(e interface{}) {
						self.LoggerMap["freeip"].Error(e)
						//self.LoggerMap["freeip"].Error(debug.Stack())
					})
				}
				if len(checkSucProxyList) > 0 {
					//初始化redis
					glib.Try(func() {
						jsons, errs := json.Marshal(checkSucProxyList) //转换成JSON返回的是byte[]
						if errs != nil {
							self.LoggerMap["freeip"].Error("json err:", errs.Error())
						}
						if string(jsons) != "[]" {
							err := redis.Set(key, string(jsons))
							if err != nil {
								self.LoggerMap["freeip"].Error(fmt.Sprintf("insert redis %s", err))
							}
							self.LoggerMap["freeip"].Info(fmt.Sprintf("insert redis %s suc...", string(jsons)))
						}
					}, func(e interface{}) {
						//self.LoggerMap["freeip"].Error(e.(error).Error())
						self.LoggerMap["freeip"].Error(e)
					})
				}

				time.Sleep(time.Second * time.Duration(5))
			}

		}
	}()

	wg.Add(1)
	go func() {
		http.ListenAndServe("0.0.0.0:54525", nil)
	}()

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
