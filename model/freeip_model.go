package model

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	glib "lib"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
	"zl_spider/config"

	l4g "code.google.com/p/log4go"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type FreeIpModel struct {
	Coding    config.UserConfigInfo
	Cfg       config.Config
	l4gLogger *l4g.Logger
	Sigs      chan os.Signal
	Running   bool
	Wg        sync.WaitGroup
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

func NewFreeIpModel(coding config.UserConfigInfo, cfg config.Config, logger *l4g.Logger, sigs chan os.Signal, wg sync.WaitGroup) *FreeIpModel {
	freeip_model := &FreeIpModel{Coding: coding, Cfg: cfg, l4gLogger: logger, Sigs: sigs, Wg: wg}
	return freeip_model
}

//获取免费代理列表
func (self *FreeIpModel) GetFreeProxyList() []FreeIpList {
	var free_ip_lists []FreeIpList
	pages := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range pages {
		url := fmt.Sprintf("%s?order_by=speed&order_rule=ASC&page=%d", self.Coding.Url, v)
		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {

			panic(err)
		}
		jsonStr := string(body)
		//fmt.Println(jsonStr)
		//jsonStr := `{"code":0,"msg":"\u6210\u529f","data":{"current_page":1,"data":[{"unique_id":"dfc97978f56f4128752303e64cba4a50","ip":"47.106.59.75","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"http","isp":"\u963f\u91cc\u4e91","speed":31,"validated_at":"2020-01-13 18:29:25","created_at":"2020-01-13 18:29:25","updated_at":"2020-01-13 18:29:30"},{"unique_id":"f5e7c94618aef75ef14b81250eae52b0","ip":"43.255.228.150","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u5e7f\u5dde","anonymity":2,"protocol":"https","isp":"\u8054\u901a","speed":37,"validated_at":"2020-01-13 18:29:39","created_at":"2020-01-12 22:45:45","updated_at":"2020-01-13 18:29:39"},{"unique_id":"5b38fd1ffa43e59e1cd6fb239feebfbb","ip":"47.112.112.102","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"http","isp":"\u963f\u91cc\u4e91","speed":39,"validated_at":"2020-01-13 18:31:52","created_at":"2020-01-13 03:22:21","updated_at":"2020-01-13 18:31:52"},{"unique_id":"e72e573cdf5f416eb2989992948c3bd1","ip":"43.255.228.150","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u5e7f\u5dde","anonymity":2,"protocol":"http","isp":"\u8054\u901a","speed":48,"validated_at":"2020-01-13 18:29:38","created_at":"2020-01-13 03:31:04","updated_at":"2020-01-13 18:29:38"},{"unique_id":"d1ecfd07dd5ae4ed18f23d7b8ea10301","ip":"210.22.5.117","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"https","isp":"\u8054\u901a","speed":48,"validated_at":"2020-01-13 18:29:55","created_at":"2020-01-11 22:37:42","updated_at":"2020-01-13 18:29:55"},{"unique_id":"126f840c95033d9c562883f9ec490f28","ip":"120.234.63.196","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"http","isp":"\u79fb\u52a8","speed":49,"validated_at":"2020-01-13 18:31:44","created_at":"2020-01-13 14:05:21","updated_at":"2020-01-13 18:31:44"},{"unique_id":"a20d691e7cb772ce4e8156eae2186a6c","ip":"47.112.112.102","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"https","isp":"\u963f\u91cc\u4e91","speed":50,"validated_at":"2020-01-13 18:29:59","created_at":"2020-01-13 00:48:14","updated_at":"2020-01-13 18:29:59"},{"unique_id":"fb4446587daa1000d58c5cc2527f2b72","ip":"118.70.144.77","port":"3128","country":"\u8d8a\u5357","ip_address":"\u8d8a\u5357 \u6cb3\u5185 XX","anonymity":2,"protocol":"https","isp":"Finance-and-Promotin","speed":60,"validated_at":"2020-01-13 18:30:43","created_at":"2020-01-13 14:32:09","updated_at":"2020-01-13 18:30:43"},{"unique_id":"c2fad51563ca86280b04d72c7db8e244","ip":"101.95.115.196","port":"80","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u4e0a\u6d77 \u4e0a\u6d77","anonymity":2,"protocol":"http","isp":"\u7535\u4fe1","speed":60,"validated_at":"2020-01-13 18:25:24","created_at":"2020-01-13 00:34:57","updated_at":"2020-01-13 18:25:24"},{"unique_id":"c2dcb569db4db065551bdcf834f2efc8","ip":"14.29.126.132","port":"80","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u5e7f\u5dde","anonymity":2,"protocol":"http","isp":"\u7535\u4fe1","speed":66,"validated_at":"2020-01-13 18:29:57","created_at":"2020-01-12 00:32:14","updated_at":"2020-01-13 18:29:57"},{"unique_id":"c2057e795738fee0d33e5188f6263ca8","ip":"47.112.218.30","port":"8000","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"https","isp":"\u963f\u91cc\u4e91","speed":74,"validated_at":"2020-01-13 18:29:07","created_at":"2020-01-13 18:29:07","updated_at":"2020-01-13 18:29:11"},{"unique_id":"c226c48d2783865c46a7f08b9b41dbad","ip":"223.199.28.12","port":"9999","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u6d77\u5357 \u6d77\u53e3","anonymity":2,"protocol":"https","isp":"\u7535\u4fe1","speed":76,"validated_at":"2020-01-13 18:26:17","created_at":"2020-01-13 18:26:17","updated_at":"2020-01-13 18:26:46"},{"unique_id":"1638f7eadfee7d2949aff8f76769c542","ip":"210.22.5.117","port":"3128","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u5e7f\u4e1c \u6df1\u5733","anonymity":2,"protocol":"http","isp":"\u8054\u901a","speed":83,"validated_at":"2020-01-13 18:29:37","created_at":"2020-01-12 01:23:00","updated_at":"2020-01-13 18:29:37"},{"unique_id":"c25a18030c234ffe77b31592485d8106","ip":"35.247.147.167","port":"3128","country":"\u65b0\u52a0\u5761","ip_address":"\u65b0\u52a0\u5761 XX XX","anonymity":2,"protocol":"http","isp":"\u8c37\u6b4c","speed":86,"validated_at":"2020-01-13 18:29:56","created_at":"2020-01-13 16:51:05","updated_at":"2020-01-13 18:29:56"},{"unique_id":"5307293f04b2714ad90b5083558f3394","ip":"106.14.203.90","port":"80","country":"\u4e2d\u56fd","ip_address":"\u4e2d\u56fd \u4e0a\u6d77 \u4e0a\u6d77","anonymity":2,"protocol":"https","isp":"\u963f\u91cc\u4e91","speed":89,"validated_at":"2020-01-13 18:31:23","created_at":"2020-01-13 14:40:40","updated_at":"2020-01-13 18:31:23"}],"first_page_url":"https:\/\/www.freeip.top\/api\/proxy_ips?page=1","from":1,"last_page":20,"last_page_url":"https:\/\/www.freeip.top\/api\/proxy_ips?page=20","next_page_url":"https:\/\/www.freeip.top\/api\/proxy_ips?page=2","path":"https:\/\/www.freeip.top\/api\/proxy_ips","per_page":15,"prev_page_url":null,"to":15,"total":293}}`
		var free_ip_list FreeIpList
		json.Unmarshal([]byte(jsonStr), &free_ip_list)
		free_ip_lists = append(free_ip_lists, free_ip_list)
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	}
	return free_ip_lists
}

func (self *FreeIpModel) Run() interface{} {
	fmt.Println("FreeIpModel")
	go func() {
		<-self.Sigs
		self.Running = true
	}()
	var query = []string{"php", "python", "go", "js"}

	var free_ip_lists []FreeIpList
	sucProxyList := make([]string, 0)

	//redis := glib.NewRedis(self.Cfg.Redis.Host, self.Cfg.Redis.Port, self.Cfg.Redis.Passwd, self.Cfg.Redis.Select)

	for {
		if self.Running {
			self.Wg.Done()
			break
		}
		glib.Try(func() {
			free_ip_lists = self.GetFreeProxyList()
		}, func(e interface{}) {
			self.l4gLogger.Error(e)
		})

		for _, v := range free_ip_lists {
			list := v.Data.Data

			ret := false
			rand.Seed(time.Now().UnixNano()) //随机干扰
			for _, v := range list {
				mt_rand := rand.Intn(3) + 1
				glib.Try(func() {
					//reqProxy := fmt.Sprintf("%s://%s:%s", v.Protocol, v.Ip, v.Port)
					reqProxy := fmt.Sprintf("http://%s:%s", v.Ip, v.Port)
					self.l4gLogger.Info(reqProxy)
					ret = self.Curl(reqProxy, fmt.Sprintf("https://www.zhipin.com/mobile/jobs.json?city=c101010100&query=%s&page=%d", query[rand.Intn(3)], mt_rand), "")
					if ret {
						if !glib.IsExistByKey(reqProxy, sucProxyList) {
							sucProxyList = append(sucProxyList, reqProxy)
						}
					}
				}, func(e interface{}) {
					self.l4gLogger.Error(e)
				})

				time.Sleep(time.Second * time.Duration(mt_rand))
			}
		}
		/*self.l4gLogger.Info(sucProxyList)
		//初始化redis
		glib.Try(func() {
			jsons, errs := json.Marshal(sucProxyList) //转换成JSON返回的是byte[]
			if errs != nil {
				self.l4gLogger.Error(errs.Error())
			}
			err := redis.Set("freeIps", string(jsons))
			if err != nil {
				self.l4gLogger.Error(fmt.Sprintf("insert redis %s", err))
			}
			self.l4gLogger.Info(string(jsons))
		}, func(e interface{}) {
			self.l4gLogger.Error(e)
		})*/

		time.Sleep(time.Second * time.Duration(rand.Intn(120)))
	}

	//redis.Close()
	return ""

}

func (self *FreeIpModel) Curl(proxy string, req_url string, checkStr string) bool {
	var tr *http.Transport
	proxyHandle, _ := url.Parse(proxy)

	tr = &http.Transport{
		TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
		Proxy:             http.ProxyURL(proxyHandle),
		DisableKeepAlives: true,
		//DisableCompression: true,
	}
	defer tr.CloseIdleConnections()

	c := &http.Client{
		Timeout:   time.Duration(10) * time.Second,
		Transport: tr,
	}

	var resp *http.Response
	var req *http.Request

	req, _ = http.NewRequest("GET", req_url, nil)
	resp, err := c.Do(req)

	if err != nil {
		if resp != nil {
			defer resp.Body.Close()
		}
		panic(err)
		//self.l4gLogger.Error(fmt.Sprintf("req: %s, err: %v", req_url, err))
		//return false
	}
	if resp.StatusCode == 200 {
		body := glib.ReadAll2(resp.Body)

		respStr := string(body)
		if !strings.Contains(respStr, "当前访问的IP存在异常行为") {
			defer resp.Body.Close()
			return true
		}

		if strings.Contains(respStr, checkStr) {
			defer resp.Body.Close()
			return true
		}
	}
	defer resp.Body.Close()
	return false
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

func (self *FreeIpModel) ProcessHttpProxyStr(ip string, port string) string {
	return fmt.Sprintf("http://%s:%s", ip, port)
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
