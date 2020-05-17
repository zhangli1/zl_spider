package model

import (
	"fmt"
	glib "lib"
	"math/rand"
	"os"
	"sync"
	"time"
	"zl_spider/config"
	"zl_spider/lib"

	"github.com/PuerkitoBio/goquery"

	l4g "code.google.com/p/log4go"
)

type KuaidailiModel struct {
	Coding    config.UserConfigInfo
	Cfg       config.Config
	l4gLogger *l4g.Logger
	Sigs      chan os.Signal
	Running   bool
	Wg        sync.WaitGroup
	XicidailiModel
	FreeIpModel
}

func NewKuaidailiModel(coding config.UserConfigInfo, cfg config.Config, logger *l4g.Logger, sigs chan os.Signal, wg sync.WaitGroup) *KuaidailiModel {
	kuaidaili_model := &KuaidailiModel{Coding: coding, Cfg: cfg, l4gLogger: logger, Sigs: sigs, Wg: wg}
	return kuaidaili_model
}

//获取免费代理列表
func (self *KuaidailiModel) GetProxyList() []XicidailiData {
	kuaidailiList := make([]XicidailiData, 0)

	pages := []int{}
	for l := 1; l <= 10; l++ {
		pages = append(pages, l)
	}
	for _, v := range pages {
		self.l4gLogger.Info(fmt.Sprintf("%d pages...", v))
		url := fmt.Sprintf("%s%d", self.Coding.Url, v)
		param := make(map[string]interface{})
		res := lib.NewRequest(self.Coding, self.Cfg).Run(url, param)
		//fmt.Println(res.Html())

		res.Find("tr").Each(func(i int, s *goquery.Selection) {
			kuaidailidata := XicidailiData{}
			s.Find("td").Each(func(ii int, ss *goquery.Selection) {
				val := ss.Text()
				if ii == 0 {
					kuaidailidata.Ip = val
				}
				if ii == 1 {
					kuaidailidata.Port = val
				}
			})
			if kuaidailidata.Ip != "" {
				//过滤
				proxy := self.FreeIpModel.ProcessHttpProxyStr(kuaidailidata.Ip, kuaidailidata.Port)
				if !glib.IsExistByKey(proxy, FilterIPList) {
					kuaidailiList = append(kuaidailiList, kuaidailidata)
				}
			}
		})
		//self.l4gLogger.Info("xicidailiList", xicidailiList)
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	}
	return kuaidailiList
}

func (self *KuaidailiModel) Run() interface{} {
	fmt.Println("KuaidailiModel")
	go func() {
		<-self.Sigs
		self.Running = true
	}()

	var ip_lists []XicidailiData
	sucProxyList := make(map[string]SucDataStruct, 0)

	for {
		if self.Running {
			self.Wg.Done()
			break
		}
		self.l4gLogger.Info("begin get GetProxyList")
		glib.Try(func() {
			ip_lists = self.GetProxyList()
		}, func(e interface{}) {
			self.l4gLogger.Error(e)
		})
		self.l4gLogger.Info("end get GetProxyList")

		sucProxyList = NewXicidailiModel(self.Coding, self.Cfg, self.l4gLogger, self.Sigs, self.Wg).checkRequest(ip_lists, 1)
		self.l4gLogger.Info("check proxy: ", sucProxyList)

	}

	return ""

}

func (self *KuaidailiModel) Destruct(param interface{}) interface{} {
	ret_data := make(map[string]interface{})
	ret_data["freeip"] = param
	return ret_data

}
