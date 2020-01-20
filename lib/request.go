package lib

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	//	"math/rand"
	"net/http"
	//"os"
	"time"
	"zl_spider/config"

	"github.com/PuerkitoBio/goquery"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Request struct {
	UserConfigInfo config.UserConfigInfo
	Cfg            config.Config
}

func NewRequest(info config.UserConfigInfo, cfg config.Config) *Request {
	request := &Request{UserConfigInfo: info, Cfg: cfg}
	return request
}

func (self *Request) Run(req_url string, req_param map[string]interface{}) *goquery.Document {
	return self.http_request(req_url, req_param)
}

func (self *Request) http_request(req_url string, req_param map[string]interface{}) *goquery.Document {
	var tr *http.Transport
	//先临时固定代理地址
	/*proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(self.Cfg.Proxy.Links)
	}*/
	if self.Cfg.Proxy.IsUse == 1 {

		tr = &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true,
			//Proxy:           proxy,
			//DisableCompression: true,
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig:   &tls.Config{InsecureSkipVerify: true},
			DisableKeepAlives: true,
			//Proxy:           proxy,
			//DisableCompression: true,
		}
	}
	timeout_s := self.UserConfigInfo.Timeout
	c := &http.Client{
		Timeout:   time.Duration(timeout_s) * time.Second,
		Transport: tr,
	}

	var resp *http.Response
	var req *http.Request

	b, _ := json.Marshal(req_param)
	body := bytes.NewBuffer(b)
	if len(req_param) < 1 {
		//resp, _ = c.Get(self.UserConfig.Url)
		req, _ = http.NewRequest("GET", req_url, nil)
	} else {
		//resp, _ = c.Post(self.UserConfig.Url, "application/x-www-form-urlencoded", body)
		req, _ = http.NewRequest("POST", req_url, body)
	}

	/*mua := []string{
		"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
		"Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
		"Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
		"Mozilla/4.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)",
		"Mozilla/3.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6",
		"Mozilla/4.1 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0",
		"Mozilla/4.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5",
		"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6",
		"Mozilla/4.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52",
	}

	rmua := mua[rand.Intn(16)]
	*/

	//req.Header.Set("User-Agent", rmua)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")
	req.Header.Set("Content-Type", "application/json")

	//添加cookie
	/*cookieStr := "lastCity=101200100; _uab_collina=156680011187464472376556; __c=1578469226; __g=-; __l=l=%2Fwww.zhipin.com%2Fweb%2Fcommon%2Fsecurity-check.html%3Fseed%3DxvsC2ObcAnjnNqlECUh3WfaSG%252Ff%252F6x47dA9eu%252BB3gwg%253D%26name%3D7ee39cf5%26ts%3D1578469225239%26callbackUrl%3D%252Fc101280600%253Fquery%253Dphp%2526page%253D1%26srcReferer%3D&r=&friend_source=0&friend_source=0; Hm_lvt_194df3105ad7148dcf2b98a91b5e727a=1577191781,1578469228; __a=66176214.1566800112.1577191780.1578469226.59.7.13.19; Hm_lpvt_194df3105ad7148dcf2b98a91b5e727a=1578642197; __zp_stoken__=887aUp5dIbX1F6IrVtZ8IY%2FpDhJtU2VztA6RctCg%2BgpAUoDGYc%2FsJWNTk8Wqo4kQnkYyCiMbsvtYXcgl2ryBBztjaxIK06Bwxf15y6ZsB86GDgpjkWJqMc3FrQt5VKX77nut"
	for _, v := range strings.Split(cookieStr, "; ") {
		cv := strings.Split(v, "=")
		cookie1 := &http.Cookie{
			Name:     cv[0],
			Value:    cv[1],
			HttpOnly: true,
		}
		req.AddCookie(cookie1)
	}*/
	resp, err := c.Do(req)
	if err != nil {
		defer resp.Body.Close()
		panic(err)
	}
	res, _ := goquery.NewDocumentFromResponse(resp)
	//fmt.Println(res)
	//os.Exit(-1)
	defer resp.Body.Close()
	return res
}

func (self *Request) StartChrome(Url string, Proxy string) string {
	opts := []selenium.ServiceOption{}
	caps := selenium.Capabilities{
		"browserName":     "chrome",
		"excludeSwitches": "enable-automation",
	}

	// 禁止加载图片，加快渲染速度
	imagCaps := map[string]interface{}{
		"profile.managed_default_content_settings.images": 2,
	}

	links := Proxy

	chromeCaps := chrome.Capabilities{
		Prefs: imagCaps,
		Path:  "",
		Args: []string{
			"--headless", // 设置Chrome无头模式
			"--no-sandbox",
			"--user-agent=Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36", // 模拟user-agent，防反爬
			fmt.Sprintf("--proxy-server=%s", links),
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
		panic(fmt.Sprintf("Error starting the ChromeDriver server: %v", err))
	}
	// 调起chrome浏览器
	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 19515))
	defer webDriver.Close()
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
	webDriver.DeleteAllCookies()
	if err == nil {
		return ret
	}
	return ""
}
