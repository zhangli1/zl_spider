package controller

import (
	"sync"
	//"fmt"
	"os"
	"zl_spider/config"
	"zl_spider/model"

	l4g "code.google.com/p/log4go"
)

type Route struct {
	Coding    config.UserConfigInfo
	Cfg       config.Config
	LoggerMap map[string]*l4g.Logger
	Sigs      chan os.Signal
	Wg        sync.WaitGroup
}

func NewRoute(coding config.UserConfigInfo, cfg config.Config, logger_map map[string]*l4g.Logger, sigs chan os.Signal, wg sync.WaitGroup) *Route {
	return &Route{Coding: coding, Cfg: cfg, LoggerMap: logger_map, Sigs: sigs, Wg: wg}
}

//所有路由集中在这里，由model提供数据
func (self *Route) Run() interface{} {
	model_list := make(map[string]model.Model, 0)
	model_list["boss"] = model.NewBossModel(self.Coding, self.Cfg, self.LoggerMap["boss"], self.Sigs, self.Wg)
	model_list["bd"] = model.NewBdModel(self.Coding, self.Cfg)
	model_list["zhihu"] = model.NewZhihuModel(self.Coding, self.Cfg)
	model_list["feixiaohao"] = model.NewFeixiaoHaoModel(self.Coding, self.Cfg)
	model_list["v2ex"] = model.NewV2exModel(self.Coding, self.Cfg)
	model_list["freeip"] = model.NewFreeIpModel(self.Coding, self.Cfg, self.LoggerMap["freeip"], self.Sigs, self.Wg)

	if model_list[self.Coding.ModelPrefix] != nil {
		return model_list[self.Coding.ModelPrefix].Run()
	}
	return nil
}
