package model

import (
	"encoding/json"
	"fmt"
	"html"
	glib "lib"
	"math/rand"
	"os"
	"runtime/debug"
	"strings"
	"sync"
	"time"
	"zl_spider/config"
	"zl_spider/lib"

	l4g "code.google.com/p/log4go"
	"github.com/PuerkitoBio/goquery"
	ghtml "golang.org/x/net/html"
)

type BossModel struct {
	Coding    config.UserConfigInfo
	Cfg       config.Config
	l4gLogger *l4g.Logger
	Sigs      chan os.Signal
	Running   bool
	Wg        sync.WaitGroup
}

type SearchField struct {
	//职位类型
	Query string
	//页数
	Page int
}

func NewBossModel(coding config.UserConfigInfo, cfg config.Config, logger *l4g.Logger, sigs chan os.Signal, wg sync.WaitGroup) *BossModel {
	boss_model := &BossModel{Coding: coding, Cfg: cfg, l4gLogger: logger, Sigs: sigs, Wg: wg}
	return boss_model
}

func (self *BossModel) Run() interface{} {
	fmt.Println("BossModel")
	city_list := make(map[string]string, 0)
	city_list["c101010100"] = "北京"
	city_list["c101020100"] = "上海"
	city_list["c101280100"] = "广州"
	city_list["c101280600"] = "深圳"
	city_list["c101200100"] = "武汉"

	position := make([]string, 0)
	position = []string{"php", "python", "golang"}

	//coding := self.Coding

	//生成1到10页码，顺序打乱
	page_list := []int{1, 5, 4, 2, 3, 6, 10, 8, 9, 7}

	req_url_tmp := self.Coding.Url

	type BossRet struct {
		HasMore bool
		Rescode int
		Html    string
	}
	// 初始化随机数的资源库, 如果不执行这行, 不管运行多少次都返回同样的值
	rand.Seed(time.Now().UnixNano())
	//优雅的退出
	go func() {
		<-self.Sigs
		self.Running = true
	}()

	redis := glib.NewRedis(self.Cfg.Redis.Host, self.Cfg.Redis.Port, self.Cfg.Redis.Passwd, self.Cfg.Redis.Select)

	for {
		if self.Running {
			self.Wg.Done()
			break
		}
		for ck, cv := range city_list {
			for _, pv := range position {
				req_url := ""

				//page := 1
				for _, page := range page_list {
					//for {
					req_url = fmt.Sprintf("%s?city=%s&query=%s&page=%d", req_url_tmp, ck, pv, page)
					self.l4gLogger.Info(req_url)

					self.Coding.Url = req_url

					req_ret := ""
					proxy := ""
					glib.Try(func() {
						//req_ret, _ := lib.NewRequest(self.Coding, self.Cfg).Run().Html()
						//req_ret = self.StartChrome(req_url)

						jsonstrByRedis, _ := redis.GET("freeIps")
						ipListStr := glib.B2S(jsonstrByRedis.([]uint8))

						var ipList []string
						json.Unmarshal([]byte(ipListStr), &ipList)
						if len(ipList) < 1 {
							proxy = self.Cfg.Proxy.Links
						} else {
							linkTmp := ipList[rand.Intn(len(ipList))]
							proxy = strings.Replace(strings.Replace(linkTmp, "http://", "", -1), "https://", "", -1)
						}

						req_ret = lib.NewRequest(self.Coding, self.Cfg).StartChrome(req_url, proxy)
					}, func(e interface{}) {
						debug.PrintStack()
						self.l4gLogger.Error(e)
						self.l4gLogger.Error(debug.Stack())
					})
					if len(req_ret) < 50 {
						continue
					}
					jsonStr := req_ret[84 : len(req_ret)-20]

					var bossret BossRet
					json.Unmarshal([]byte(jsonStr), &bossret)
					htmlData := html.UnescapeString(bossret.Html)

					htmlNode, _ := ghtml.Parse(strings.NewReader(htmlData))
					ret := self.Parse(goquery.NewDocumentFromNode(htmlNode), cv, page)

					if len(ret) < 1 {
						continue
					}

					//写入els数据
					for _, rv := range ret {
						self.l4gLogger.Info(rv)
						glib.Try(func() {
							self.WriteEls(rv)
						}, func(e interface{}) {
							debug.PrintStack()
							fmt.Println(e)
							self.l4gLogger.Error(debug.Stack())
						})
					}
					time.Sleep(time.Second * time.Duration(rand.Intn(5)))
					//page++
				}
			}
			time.Sleep(time.Second * time.Duration(rand.Intn(self.Cfg.Frequency.City)))
		}
		time.Sleep(time.Second * time.Duration(rand.Intn(300)))
	}

	return ""

}

//写入Elasticsearch

func (self *BossModel) WriteEls(param lib.JobData) {
	coding := self.Coding
	//self.Coding.Url = self.Cfg.Elasticsearch.Url
	c_param := glib.Struct2Map(param)
	//self.Coding.Param = c_param
	json_ret := lib.NewRequest(self.Coding, self.Cfg).Run(self.Cfg.Elasticsearch.Url, c_param).Text()

	var mapResult map[string]interface{}
	if err := json.Unmarshal([]byte(json_ret), &mapResult); err != nil {
		panic(err)
	}
	if _, ok := mapResult["_id"]; !ok {
		self.l4gLogger.Info(mapResult["status"])
	}
	self.Coding = coding
}

//解析html node
func (self *BossModel) Parse(resp *goquery.Document, city string, page int) []lib.JobData {
	var line lib.JobData
	job_data := make([]lib.JobData, 0)

	self.l4gLogger.Info(resp.Html())

	//return job_data

	resp.Find("li").Each(func(i int, s *goquery.Selection) {
		line = lib.JobData{}
		//抓取网站
		line.WebSite = self.Coding.ModelPrefix
		//城市
		line.City = city

		//职位信息
		line.JobTitle = s.Find(".title h4").Text()
		//薪水
		line.Salary = s.Find(".salary").Text()
		//地址、经验、学历
		p_add_data := make([]string, 0)
		s.Find(".msg em").Each(func(i int, s *goquery.Selection) {
			p_add_data = append(p_add_data, s.Text())
		})

		if len(p_add_data) > 0 {
			line.Address = p_add_data[0]
			if len(p_add_data) > 2 {
				line.Empirical = p_add_data[1]
				line.Education = p_add_data[2]
			} else if len(p_add_data) > 1 && (strings.Contains(p_add_data[1], "-") || strings.Contains(p_add_data[1], "年") || strings.Contains(p_add_data[1], "经验不限")) {
				line.Education = p_add_data[1]
			}
		}

		//公司名
		line.CompanyName = s.Find(".name").Text()
		//公司类型
		/*c_company_data, _ := s.Find(".job-primary .info-company .company-text p").Html()
		vline_data := strings.Split(c_company_data, "<em class=\"vline\"></em>")
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
		*/
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
