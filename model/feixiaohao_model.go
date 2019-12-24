package model

import (
	//"encoding/json"
	//"fmt"
	//glib "lib"
	//"math/rand"
	//	"os"
	//"regexp"
	//	"strings"
	//	"time"
	"zl_spider/config"
	//	"zl_spider/lib"
	//	"github.com/PuerkitoBio/goquery"
)

type FeixiaoHaoModel struct {
	Coding config.UserConfigInfo
	Cfg    config.Config
}

func NewFeixiaoHaoModel(coding config.UserConfigInfo, cfg config.Config) *FeixiaoHaoModel {
	boss_model := &FeixiaoHaoModel{Coding: coding, Cfg: cfg}
	return boss_model
}

func (self *FeixiaoHaoModel) Run() interface{} {
	return "Feixiaohao"
}

func (self *FeixiaoHaoModel) Destruct(param interface{}) interface{} {
	ret_data := make(map[string]interface{})
	ret_data["Feixiaohao"] = param
	return ret_data
}
