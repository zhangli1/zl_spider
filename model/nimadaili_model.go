package model

import (
	"fmt"
	glib "lib"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
	"zl_spider/config"
	"zl_spider/lib"

	"github.com/PuerkitoBio/goquery"

	l4g "code.google.com/p/log4go"
)

type NimadailiModel struct {
	Coding    config.UserConfigInfo
	Cfg       config.Config
	l4gLogger *l4g.Logger
	Sigs      chan os.Signal
	Running   bool
	Wg        sync.WaitGroup
	XicidailiModel
	FreeIpModel
}

func NewNimadailiModel(coding config.UserConfigInfo, cfg config.Config, logger *l4g.Logger, sigs chan os.Signal, wg sync.WaitGroup) *NimadailiModel {
	nimadaili_model := &NimadailiModel{Coding: coding, Cfg: cfg, l4gLogger: logger, Sigs: sigs, Wg: wg}
	return nimadaili_model
}

//获取免费代理列表
func (self *NimadailiModel) GetProxyList() []XicidailiData {
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
					line := strings.Split(val, ":")
					kuaidailidata.Ip = line[0]
					kuaidailidata.Port = line[1]
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

func (self *NimadailiModel) Run() interface{} {
	fmt.Println("NimadailiModel")
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

func (self *NimadailiModel) Destruct(param interface{}) interface{} {
	ret_data := make(map[string]interface{})
	ret_data["freeip"] = param
	return ret_data

}
