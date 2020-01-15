package model

import (

	//"encoding/json"
	"fmt"
	//glib "lib"
	//"math/rand"
	//	"os"
	//"regexp"
	//	"strings"
	//	"time"
	"zl_spider/config"
	"zl_spider/lib"

	"github.com/PuerkitoBio/goquery"
)

type FeixiaoHaoModel struct {
	Coding config.UserConfigInfo
	Cfg    config.Config
}

func NewFeixiaoHaoModel(coding config.UserConfigInfo, cfg config.Config) *FeixiaoHaoModel {
	feixiaohao_model := &FeixiaoHaoModel{Coding: coding, Cfg: cfg}
	return feixiaohao_model
}

func (self *FeixiaoHaoModel) Run() interface{} {
	var req_param map[string]interface{}
	ret := lib.NewRequest(self.Coding, self.Cfg).Run(self.Coding.Url, req_param)
	return self.Parse(ret)
}

func (self *FeixiaoHaoModel) Parse(doc *goquery.Document) interface{} {
	fmt.Println(doc.Find(".ivu-table-body").Find("").Html())
	doc.Find(".ivu-table-body .ivu-table-row").Each(func(i int, s *goquery.Selection) {
		//fmt.Println(s.Html())
	})
	return nil
}

func (self *FeixiaoHaoModel) Destruct(param interface{}) interface{} {
	ret_data := make(map[string]interface{})
	ret_data["Feixiaohao"] = param
	return ret_data
}
