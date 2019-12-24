package model

import (
	"encoding/json"
	"fmt"
	glib "lib"
	"math/rand"
	//	"os"
	"regexp"
	"strings"
	"time"
	"zl_spider/config"
	"zl_spider/lib"

	"github.com/PuerkitoBio/goquery"
)

type BossModel struct {
	Coding config.UserConfigInfo
	Cfg    config.Config
}

type SearchField struct {
	//职位类型
	Query string
	//页数
	Page int
}

func NewBossModel(coding config.UserConfigInfo, cfg config.Config) *BossModel {
	boss_model := &BossModel{Coding: coding, Cfg: cfg}
	return boss_model
}

func (self *BossModel) Run() interface{} {
	city_list := make(map[string]string, 0)
	//city_list["c101010100"] = "北京"
	//city_list["c101020100"] = "上海"
	//city_list["c101280100"] = "广州"
	city_list["c101280600"] = "深圳"
	city_list["c101200100"] = "武汉"

	position := make([]string, 0)
	//position = []string{"php", "python", "golang"}
	position = []string{"php", "golang"}
	//position = []string{"python", "golang"}

	coding := self.Coding

	//生成1到10页码，顺序打乱
	page_list := []int{1, 5, 4, 2, 3}

	requ_url_tmp := self.Coding.Url
	for ck, cv := range city_list {
		for _, pv := range position {
			req_url := ""

			//page := 1
			for _, page := range page_list {
				//for {
				req_url = fmt.Sprintf("%s/%s?query=%s&page=%d", requ_url_tmp, ck, pv, page)
				fmt.Println(req_url)
				self.Coding.Url = req_url
				ret := self.Parse(lib.NewRequest(self.Coding, self.Cfg).Run(), cv, page)
				if len(ret) < 1 {
					//break
					continue
				}

				//写入els数据
				for _, rv := range ret {
					self.WriteEls(rv)
				}
				time.Sleep(time.Second * time.Duration(rand.Intn(5)))
				page++
			}
		}
		time.Sleep(time.Second * time.Duration(self.Cfg.Frequency.City))
	}

	self.Coding = coding
	//self.Resp = lib.NewRequest(self.Coding).Run()

	//return self.Destruct(job_data)
	return ""

}

//写入Elasticsearch

func (self *BossModel) WriteEls(param lib.JobData) {
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
func (self *BossModel) Parse(resp *goquery.Document, city string, page int) []lib.JobData {
	var line lib.JobData
	job_data := make([]lib.JobData, 0)

	fmt.Println(resp.Html())
	//获取当前真实的页码数
	p := glib.StringToInt(resp.Find(".page .cur").Text())
	if p < page {
		return job_data
	}

	resp.Find(".job-list ul li").Each(func(i int, s *goquery.Selection) {
		line = lib.JobData{}
		//抓取网站
		line.WebSite = self.Coding.ModelPrefix
		//城市
		line.City = city

		//职位信息
		line.JobTitle = s.Find(".job-primary .info-primary h3 .job-title").Text()
		//薪水
		line.Salary = s.Find(".job-primary .info-primary h3 .red").Text()
		//地址、经验、学历
		p_add_data, _ := s.Find(".job-primary .info-primary p").Html()

		vline_data := strings.Split(p_add_data, "<em class=\"vline\"></em>")

		line.Address = vline_data[0]
		if len(vline_data) > 2 {
			line.Empirical = vline_data[1]
			line.Education = vline_data[2]
		} else if len(vline_data) > 1 && (strings.Contains(vline_data[1], "-") || strings.Contains(vline_data[1], "年") || strings.Contains(vline_data[1], "经验不限")) {
			line.Education = vline_data[1]
		}

		//公司名
		line.CompanyName = s.Find(".job-primary .info-company .company-text h3").Text()
		//公司类型
		c_company_data, _ := s.Find(".job-primary .info-company .company-text p").Html()
		vline_data = strings.Split(c_company_data, "<em class=\"vline\"></em>")
		line.CompanyType = vline_data[0]
		//公司人数
		if len(vline_data) > 2 {
			line.FinancingSituation = vline_data[1]
			line.Person = vline_data[2]
		} else if len(vline_data) > 1 && (strings.Contains(vline_data[1], "-") || strings.Contains(vline_data[1], "人")) {
			line.Person = vline_data[1]
		}
		//招骋人和招骋人title
		r_person_data, _ := s.Find(".info-publis h3").Html()
		vline_data2 := strings.Split(r_person_data, "<em class=\"vline\"></em>")

		reg, _ := regexp.Compile("<.*>")
		line.RecruitName = reg.ReplaceAllString(vline_data2[0], "")

		if len(vline_data2) > 1 {
			line.RecruitPosition = vline_data2[1]
		}

		//更新时间
		line.UpdateTime = self.ProcessTime(s.Find(".info-publis p").Text())
		line.CreateTime = glib.TimestampToDate("", glib.GetCurrentTime())
		job_data = append(job_data, line)
	})
	return job_data

	/*if len(job_data) > 0 {
		jsons, _ := json.Marshal(job_data)
		redis := glib.NewRedis(self.Cfg.Redis.Host, self.Cfg.Redis.Port, self.Cfg.Redis.Passwd, self.Cfg.Redis.Select)
		err := redis.Set("boss", string(jsons))
		if err == nil {
			fmt.Println("suc")
		}
	}*/
}

func (self *BossModel) Destruct(param interface{}) interface{} {
	ret_data := make(map[string]interface{})
	ret_data["boss"] = param
	return ret_data

}

//解析更新时间
func (self *BossModel) ProcessTime(UpdateTimeString string) string {
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
