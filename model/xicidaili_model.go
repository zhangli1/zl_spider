package model

import (
	"fmt"
	glib "lib"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"
	"zl_spider/config"
	"zl_spider/lib"

	"github.com/PuerkitoBio/goquery"

	l4g "code.google.com/p/log4go"
)

type XicidailiModel struct {
	Coding    config.UserConfigInfo
	Cfg       config.Config
	l4gLogger *l4g.Logger
	Sigs      chan os.Signal
	Running   bool
	Wg        sync.WaitGroup
	FreeIpModel
}

type SucDataStruct struct {
	Url      string
	List     []string
	CheckStr string
}

var FilterIPList = make([]string, 0)

var SucData = make(map[string]SucDataStruct, 0)

var ProxyTableData []lib.Proxy

func NewXicidailiModel(coding config.UserConfigInfo, cfg config.Config, logger *l4g.Logger, sigs chan os.Signal, wg sync.WaitGroup) *XicidailiModel {
	xicidaili_model := &XicidailiModel{Coding: coding, Cfg: cfg, l4gLogger: logger, Sigs: sigs, Wg: wg}
	return xicidaili_model
}

type XicidailiData struct {
	Ip   string
	Port string
}

//获取免费代理列表
func (self *XicidailiModel) GetProxyList() []XicidailiData {
	xicidailiList := make([]XicidailiData, 0)

	pages := []int{}
	for l := 1; l <= 10; l++ {
		pages = append(pages, l)
	}
	for _, v := range pages {
		self.l4gLogger.Info(fmt.Sprintf("%d pages...", v))
		url := fmt.Sprintf("%s%d", self.Coding.Url, v)
		param := make(map[string]interface{})
		res := lib.NewRequest(self.Coding, self.Cfg).Run(url, param)
		//fmt.Println(res.Html())

		res.Find("tr").Each(func(i int, s *goquery.Selection) {
			xicidailidata := XicidailiData{}
			s.Find("td").Each(func(ii int, ss *goquery.Selection) {
				if ii == 0 {
					val := ss.Text()
					line := strings.Split(val, ":")
					xicidailidata.Ip = line[0]
					xicidailidata.Port = line[1]
				}
			})
			if xicidailidata.Ip != "" {
				//过滤
				proxy := self.FreeIpModel.ProcessHttpProxyStr(xicidailidata.Ip, xicidailidata.Port)
				if !glib.IsExistByKey(proxy, FilterIPList) {
					xicidailiList = append(xicidailiList, xicidailidata)
				}
			}
		})
		//self.l4gLogger.Info("xicidailiList", xicidailiList)
		time.Sleep(time.Second * time.Duration(rand.Intn(5)))
	}
	return xicidailiList
}

func (self *XicidailiModel) checkRequest(ip_lists []XicidailiData, num int) map[string]SucDataStruct {
	var SucOneProxyList = make([]string, 0)

	ProxyTableData = lib.NewDb(self.l4gLogger, self.Cfg.MYSQL).GetProxyData("select id,url,template,paramList,list,checkStr from proxy where switch = 1")

	length := 200
	ch := make(chan bool, length)

	if len(ProxyTableData) > 0 {
		for _, v := range ProxyTableData {
			var paramList map[string]interface{}

			paramList = glib.JsonToMap(v.ParamList)

			//最短列表
			shortListLength := 0

			//生成链接
			if len(paramList) < 1 {
				continue
			}

			list := strings.Split(v.Template, "&")
			if len(list) < 1 {
				continue
			}

			listData := make(map[string][]string, 0)

			for _, itemv := range list {
				temp := strings.Split(itemv, "=")
				if paramList[temp[0]] != nil {
					for _, ptv := range paramList[temp[0]].([]interface{}) {
						if shortListLength == 0 {
							shortListLength = len(paramList[temp[0]].([]interface{}))
						} else {
							if shortListLength > 0 && len(paramList[temp[0]].([]interface{})) < shortListLength {
								shortListLength = len(paramList[temp[0]].([]interface{}))
							}
						}
						if !glib.IsExistByKey(temp[0], listData) {
							listData[temp[0]] = make([]string, 0)
						}
						if _, ok := ptv.(float64); ok {
							ptv = glib.Float64ToString(ptv.(float64))
						}
						listData[temp[0]] = append(listData[temp[0]], fmt.Sprintf(itemv, ptv))
					}
				}
			}

			//组合成新的参数
			var urlList []string
			if len(listData) > 0 {
				for i := 0; i < shortListLength; i++ {
					var url []string
					for _, lv := range listData {
						url = append(url, lv[i])
					}
					urlList = append(urlList, fmt.Sprintf("%s?%s", v.Url, strings.Join(url, "&")))
				}
			}

			if len(urlList) < 1 {
				continue
			}

			id := fmt.Sprintf("%d", v.ID)

			var testUrl string
			for _, execUrl := range urlList {
				testUrl = execUrl
				//第一次请求
				for _, iv := range ip_lists {
					ret := false
					if len(ch) >= length {
						time.Sleep(time.Second * time.Duration(1))
					}

					reqProxy := self.FreeIpModel.ProcessHttpProxyStr(iv.Ip, iv.Port)
					go glib.Try(func() {
						ch <- true
						ret = self.FreeIpModel.Curl(reqProxy, execUrl, v.CheckStr)
						self.l4gLogger.Info(ret, num, reqProxy)
						if ret {
							if !glib.IsExistByKey(reqProxy, SucOneProxyList) && !glib.IsExistByKey(reqProxy, FilterIPList) {
								SucOneProxyList = append(SucOneProxyList, reqProxy)
							}
							self.l4gLogger.Info("sucOneProxyList2", SucOneProxyList)
						} else {
							if !glib.IsExistByKey(reqProxy, FilterIPList) {
								FilterIPList = append(FilterIPList, reqProxy)
							}
						}
						<-ch
					}, func(e interface{}) {
						self.l4gLogger.Info(num, reqProxy)
						self.l4gLogger.Error(num, e)
						<-ch
					})
				}
			}

			var sucdata SucDataStruct
			if _, ok := SucData[id]; !ok {
				sucdata.List = make([]string, 0)
			}
			sucdata.Url = testUrl
			sucdata.List = SucOneProxyList
			sucdata.CheckStr = v.CheckStr
			SucData[id] = sucdata

		}
	}

	if len(ch) < 1 {
		time.Sleep(time.Second * time.Duration(3))
		if len(ch) < 1 {
			defer close(ch)
		}
	}

	return SucData
}

func (self *XicidailiModel) Run() interface{} {
	fmt.Println("XicidailiModel")
	go func() {
		<-self.Sigs
		self.Running = true
	}()

	var ip_lists []XicidailiData
	sucProxyList := make(map[string]SucDataStruct, 0)

	for {
		if self.Running {
			self.Wg.Done()
			break
		}
		self.l4gLogger.Info("begin get GetProxyList")
		glib.Try(func() {
			ip_lists = self.GetProxyList()
		}, func(e interface{}) {
			self.l4gLogger.Error(e)
		})
		self.l4gLogger.Info("end get GetProxyList")

		sucProxyList = self.checkRequest(ip_lists, 1)
		self.l4gLogger.Info("check proxy: ", sucProxyList)
	}

	return ""

}

/*func (self *XicidailiModel) Curl(proxy string, req_url string) bool {
	var tr *http.Transport
	proxyHandle, _ := url.Parse(proxy)

	tr = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		Proxy:           http.ProxyURL(proxyHandle),
		//DisableCompression: true,
	}

	c := &http.Client{
		Timeout:   time.Duration(5) * time.Second,
		Transport: tr,
	}

	var resp *http.Response
	var req *http.Request

	req, _ = http.NewRequest("GET", req_url, nil)
	resp, err := c.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode == 200 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		respStr := string(body)
		if !strings.Contains(respStr, "当前访问的IP存在异常行为") {
			defer resp.Body.Close()
			return true
		}
	}
	defer resp.Body.Close()
	return false
}

func (self *XicidailiModel) StartChrome(Url string) string {
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

/*chromeCaps := chrome.Capabilities{
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
/*err = webDriver.Get(Url)
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
}*/

func (self *XicidailiModel) Destruct(param interface{}) interface{} {
	ret_data := make(map[string]interface{})
	ret_data["freeip"] = param
	return ret_data

}

//解析更新时间
func (self *XicidailiModel) ProcessTime(UpdateTimeString string) string {
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
