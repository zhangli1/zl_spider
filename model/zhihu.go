package model

import (
	"encoding/json"
	"fmt"
	glib "lib"
	//"math/rand"
	"os"
	//	"regexp"
	"strings"
	"time"
	"zl_spider/config"
	"zl_spider/lib"

	"github.com/PuerkitoBio/goquery"
)

type ZhihuModel struct {
	Coding config.UserConfigInfo
	Cfg    config.Config
}

func NewZhihuModel(coding config.UserConfigInfo, cfg config.Config) *ZhihuModel {
	zhihu_model := &ZhihuModel{Coding: coding, Cfg: cfg}
	return zhihu_model
}

func (self *ZhihuModel) Run() interface{} {
	fmt.Println(self.Coding.Url)

	self.Parse(lib.NewRequest(self.Coding, self.Cfg).Run())

	os.Exit(-1)
	return ""

}

//写入Elasticsearch

func (self *ZhihuModel) WriteEls(param lib.JobData) {
	coding := self.Coding
	//self.Coding.Url = "http://els.zhangli0712.cn/boss/php"
	self.Coding.Url = self.Cfg.Elasticsearch.Url
	c_param := glib.Struct2Map(param)
	self.Coding.Param = c_param
	json_ret := lib.NewRequest(self.Coding, self.Cfg).Run().Text()

	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(json_ret), &mapResult); err != nil {
		fmt.Println(err)
	}
	if _, ok := mapResult["_id"]; !ok {
		fmt.Println(mapResult["status"])
	}
	self.Coding = coding
}

//解析html node
func (self *ZhihuModel) Parse(resp *goquery.Document) {
	//var line lib.JobData
	//job_data := make([]lib.JobData, 0)

	//fmt.Println(resp.Html())

	jsonstr := resp.Find("#js-initialData").Text()

	var ret_map map[string]interface{}
	_ = json.Unmarshal([]byte(jsonstr), &ret_map)
	answers := ret_map["initialState"].(map[string]interface{})["entities"].(map[string]interface{})["answers"]

	for k, v := range answers.(map[string]interface{}) {
		content := (strings.Replace(strings.Replace(v.(map[string]interface{})["content"].(string), "<p>", "", -1), "</p>", "", -1))
		glib.WriteFile("aaa", fmt.Sprintf("%s,%s\r\n", k, content))
	}

	/*resp.Find(".List .SearchResult-Card .AnswerItem").Each(func(i int, s *goquery.Selection) {
		fmt.Println(s.Find(".Highlight").Html())
		fmt.Println(s.Find(".RichContent .RichContent-inner").Html())
		fmt.Println("-----")
	})*/

}

func (self *ZhihuModel) Destruct(param interface{}) interface{} {
	ret_data := make(map[string]interface{})
	ret_data["boss"] = param
	return ret_data

}

//解析更新时间
func (self *ZhihuModel) ProcessTime(UpdateTimeString string) string {
	var format_date string
	UpdateTimeString = strings.Replace(UpdateTimeString, "发布于", "", -1)
	if strings.Contains(UpdateTimeString, "月") {
		format_date = strings.Replace(UpdateTimeString, "月", "-", -1)
		year, _, _ := time.Now().Date()
		format_date = fmt.Sprintf("%d-%s 00:00:00", year, strings.Replace(format_date, "日", "", -1))
	} else {
		format_date = fmt.Sprintf("%s %s:00", glib.TimestampToDate("2006-01-02", glib.GetCurrentTime()), UpdateTimeString)
	}
	return format_date
}
