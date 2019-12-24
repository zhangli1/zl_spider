package controller

import (
	//"fmt"
	"zl_spider/config"
	"zl_spider/model"
)

type Route struct {
	Coding config.UserConfigInfo
	Cfg    config.Config
}

func NewRoute(coding config.UserConfigInfo, cfg config.Config) *Route {
	return &Route{Coding: coding, Cfg: cfg}
}

func (self *Route) Run() interface{} {
	model_list := make(map[string]model.Model, 0)
	//model_list["boss"] = model.NewBossModel(self.Coding, self.Cfg)
	model_list["feixiaohao"] = model.NewFeixiaoHaoModel(self.Coding, self.Cfg)
	//model_list["bd"] = model.NewBdModel(self.Coding, self.Cfg)
	//model_list["zhihu"] = model.NewZhihuModel(self.Coding, self.Cfg)

	if model_list[self.Coding.ModelPrefix] != nil {
		return model_list[self.Coding.ModelPrefix].Run()
	}
	return nil
}
