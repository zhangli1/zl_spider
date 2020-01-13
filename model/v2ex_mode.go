package model

import (
	//"encoding/json"
	"fmt"
	glib "lib"
	//"math/rand"
	//	"os"
	//"regexp"
	//	"strings"
	//	"time"
	"zl_spider/config"
	"zl_spider/lib"

	"github.com/PuerkitoBio/goquery"
)

type V2exModel struct {
	Coding config.UserConfigInfo
	Cfg    config.Config
}

func NewV2exModel(coding config.UserConfigInfo, cfg config.Config) *V2exModel {
	v2ex_model := &V2exModel{Coding: coding, Cfg: cfg}
	return v2ex_model
}

func (self *V2exModel) Run() interface{} {
	ret := lib.NewRequest(self.Coding, self.Cfg).Run()
	return self.Parse(ret)
}

func (self *V2exModel) Parse(doc *goquery.Document) interface{} {
	fmt.Println(self.Coding.Url)

	data := map[string]int{}
	doc.Find("table").Each(func(i int, s *goquery.Selection) {
		if s.Find(".topic-link").Text() != "" {
			Title := s.Find(".topic-link").Text()
			Num := s.Find(".count_livid").Text()
			if Num == "" {
				Num = "0"
			}
			data[Title] = glib.StringToInt(Num)
		}
	})

	newData := glib.SortMapByValue(data)
	for _, v := range newData {
		fmt.Println(v)
	}

	return nil
}

func (self *V2exModel) Destruct(param interface{}) interface{} {
	ret_data := make(map[string]interface{})
	ret_data["v2ex"] = param
	return ret_data
}
