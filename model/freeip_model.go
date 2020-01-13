package model

import (
	"encoding/json"
	"fmt"
	glib "lib"
	"log"
	"strings"
	"time"
	"zl_spider/config"
	"zl_spider/lib"

	"github.com/PuerkitoBio/goquery"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type FreeIpModel struct {
	Coding config.UserConfigInfo
	Cfg    config.Config
}

type FreeIpDetail struct {
	UniqueID    string `json:"unique_id"`
	Ip          string `json:"ip"`
	Port        string `json:"port"`
	IPAddress   string `json:"ip_address"`
	Anonymity   int    `json:"anonymity"`
	Protocol    string `json:"protocol"`
	Isp         string `json:"isp"`
	Speed       int    `json:"speed"`
	ValidatedAt string `json:"validated_at"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type FreeIpList struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		CurrentPage int `json:"current_page"`
		Data        []FreeIpDetail
		LastPage    int `json:"last_page"`
		PerPage     int `json:"per_page"`
		To          int `json:"to"`
		Total       int `json:"total"`
	}
}

func NewFreeIpModel(coding config.UserConfigInfo, cfg config.Config) *FreeIpModel {
	freeip_model := &FreeIpModel{Coding: coding, Cfg: cfg}
	return freeip_model
}

func (self *FreeIpModel) Run() interface{} {
	/*url := fmt.Sprintf("%s?order_by=speed&order_rule=ASC", self.Coding.Url)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	jsonStr := string(body) */
	jsonStr := `{"code":0,"msg":"\u6210\u529f","data":{"current_page":1,"data":[{"unique_id":"dfc97978f56f4128752303e64cba4a50","ip":"47.106.59.75","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"http","isp":"\u963f\u91cc\u4e91","speed":31,"validated_at":"2020-01-13 18:29:25","created_at":"2020-01-13 18:29:25","updated_at":"2020-01-13 18:29:30"},{"unique_id":"f5e7c94618aef75ef14b81250eae52b0","ip":"43.255.228.150","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u5e7f\u5dde","anonymity":2,"protocol":"https","isp":"\u8054\u901a","speed":37,"validated_at":"2020-01-13 18:29:39","created_at":"2020-01-12 22:45:45","updated_at":"2020-01-13 18:29:39"},{"unique_id":"5b38fd1ffa43e59e1cd6fb239feebfbb","ip":"47.112.112.102","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"http","isp":"\u963f\u91cc\u4e91","speed":39,"validated_at":"2020-01-13 18:31:52","created_at":"2020-01-13 03:22:21","updated_at":"2020-01-13 18:31:52"},{"unique_id":"e72e573cdf5f416eb2989992948c3bd1","ip":"43.255.228.150","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u5e7f\u5dde","anonymity":2,"protocol":"http","isp":"\u8054\u901a","speed":48,"validated_at":"2020-01-13 18:29:38","created_at":"2020-01-13 03:31:04","updated_at":"2020-01-13 18:29:38"},{"unique_id":"d1ecfd07dd5ae4ed18f23d7b8ea10301","ip":"210.22.5.117","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"https","isp":"\u8054\u901a","speed":48,"validated_at":"2020-01-13 18:29:55","created_at":"2020-01-11 22:37:42","updated_at":"2020-01-13 18:29:55"},{"unique_id":"126f840c95033d9c562883f9ec490f28","ip":"120.234.63.196","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"http","isp":"\u79fb\u52a8","speed":49,"validated_at":"2020-01-13 18:31:44","created_at":"2020-01-13 14:05:21","updated_at":"2020-01-13 18:31:44"},{"unique_id":"a20d691e7cb772ce4e8156eae2186a6c","ip":"47.112.112.102","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"https","isp":"\u963f\u91cc\u4e91","speed":50,"validated_at":"2020-01-13 18:29:59","created_at":"2020-01-13 00:48:14","updated_at":"2020-01-13 18:29:59"},{"unique_id":"fb4446587daa1000d58c5cc2527f2b72","ip":"118.70.144.77","port":"3128","country":"\u8d8a\u5357","ip_address":"\u8d8a\u5357 \u6cb3\u5185 XX","anonymity":2,"protocol":"https","isp":"Finance-and-Promotin","speed":60,"validated_at":"2020-01-13 18:30:43","created_at":"2020-01-13 14:32:09","updated_at":"2020-01-13 18:30:43"},{"unique_id":"c2fad51563ca86280b04d72c7db8e244","ip":"101.95.115.196","port":"80","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u4e0a\u6d77 \u4e0a\u6d77","anonymity":2,"protocol":"http","isp":"\u7535\u4fe1","speed":60,"validated_at":"2020-01-13 18:25:24","created_at":"2020-01-13 00:34:57","updated_at":"2020-01-13 18:25:24"},{"unique_id":"c2dcb569db4db065551bdcf834f2efc8","ip":"14.29.126.132","port":"80","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u5e7f\u5dde","anonymity":2,"protocol":"http","isp":"\u7535\u4fe1","speed":66,"validated_at":"2020-01-13 18:29:57","created_at":"2020-01-12 00:32:14","updated_at":"2020-01-13 18:29:57"},{"unique_id":"c2057e795738fee0d33e5188f6263ca8","ip":"47.112.218.30","port":"8000","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"https","isp":"\u963f\u91cc\u4e91","speed":74,"validated_at":"2020-01-13 18:29:07","created_at":"2020-01-13 18:29:07","updated_at":"2020-01-13 18:29:11"},{"unique_id":"c226c48d2783865c46a7f08b9b41dbad","ip":"223.199.28.12","port":"9999","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u6d77\u5357 \u6d77\u53e3","anonymity":2,"protocol":"https","isp":"\u7535\u4fe1","speed":76,"validated_at":"2020-01-13 18:26:17","created_at":"2020-01-13 18:26:17","updated_at":"2020-01-13 18:26:46"},{"unique_id":"1638f7eadfee7d2949aff8f76769c542","ip":"210.22.5.117","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"http","isp":"\u8054\u901a","speed":83,"validated_at":"2020-01-13 18:29:37","created_at":"2020-01-12 01:23:00","updated_at":"2020-01-13 18:29:37"},{"unique_id":"c25a18030c234ffe77b31592485d8106","ip":"35.247.147.167","port":"3128","country":"\u65b0\u52a0\u5761","ip_address":"\u65b0\u52a0\u5761 XX XX","anonymity":2,"protocol":"http","isp":"\u8c37\u6b4c","speed":86,"validated_at":"2020-01-13 18:29:56","created_at":"2020-01-13 16:51:05","updated_at":"2020-01-13 18:29:56"},{"unique_id":"5307293f04b2714ad90b5083558f3394","ip":"106.14.203.90","port":"80","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u4e0a\u6d77 \u4e0a\u6d77","anonymity":2,"protocol":"https","isp":"\u963f\u91cc\u4e91","speed":89,"validated_at":"2020-01-13 18:31:23","created_at":"2020-01-13 14:40:40","updated_at":"2020-01-13 18:31:23"}],"first_page_url":"https:\/\/www.freeip.top\/api\/proxy_ips?page=1","from":1,"last_page":20,"last_page_url":"https:\/\/www.freeip.top\/api\/proxy_ips?page=20","next_page_url":"https:\/\/www.freeip.top\/api\/proxy_ips?page=2","path":"https:\/\/www.freeip.top\/api\/proxy_ips","per_page":15,"prev_page_url":null,"to":15,"total":293}}`
	var free_ip_list FreeIpList
	json.Unmarshal([]byte(jsonStr), &free_ip_list)
	fmt.Println(free_ip_list)

	return ""

}

func (self *FreeIpModel) StartChrome(Url string) string {
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName":     "chrome",
		"excludeSwitches": "enable-automation",
	}

	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	/*proxyStr := ""
	if self.Cfg.Proxy.Links != "" {
		proxyStr = fmt.Sprintf("--proxy-server=%s", self.Cfg.Proxy.Links)
	}*/

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless", // 设置Chrome无头模式
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36", // 模拟user-agent，防反爬
			fmt.Sprintf("--proxy-server=%s", self.Cfg.Proxy.Links),
		},
		ExcludeSwitches: []string{
			"--excludeSwitches=enable-automation",
		},
	}
	caps.AddChrome(chromeCaps)
	// 启动chromedriver，端口号可自定义
	//service, err := selenium.NewChromeDriverService("chromedriver", 19515, opts...)
	_, err := selenium.NewChromeDriverService("chromedriver", 19515, opts...)
	if err != nil {
		log.Printf("Error starting the ChromeDriver server: %v", err)
	}
	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 19515))
	if err != nil {
		panic(err)
	}
	// 这是目标网站留下的坑，不加这个在linux系统中会显示手机网页，每个网站的策略不一样，需要区别处理。
	/*webDriver.AddCookie(&selenium.Cookie{
		Name:  "cookie",
		Value: "",
	})*/

	// 导航到目标网站
	err = webDriver.Get(Url)
	if err != nil {
		panic(fmt.Sprintf("Failed to load page: %s\n", err))
	}
	//fmt.Println(webDriver.Title())
	ret, err := webDriver.PageSource()
	webDriver.Close()
	if err == nil {
		return ret
	}
	return ""
}

//写入Elasticsearch

func (self *FreeIpModel) WriteEls(param lib.JobData) {
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
func (self *FreeIpModel) Parse(resp *goquery.Document, city string, page int) []lib.JobData {
	var line lib.JobData
	job_data := make([]lib.JobData, 0)

	fmt.Println(resp.Html())

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

func (self *FreeIpModel) Destruct(param interface{}) interface{} {
	ret_data := make(map[string]interface{})
	ret_data["freeip"] = param
	return ret_data

}

//解析更新时间
func (self *FreeIpModel) ProcessTime(UpdateTimeString string) string {
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
