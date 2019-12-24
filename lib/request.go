package lib

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	//	"math/rand"
	"net/http"
	"net/url"
	//"os"
	"time"
	"zl_spider/config"

	"github.com/PuerkitoBio/goquery"
)

type Request struct {
	UserConfigInfo config.UserConfigInfo
	Cfg            config.Config
}

func NewRequest(info config.UserConfigInfo, cfg config.Config) *Request {
	request := &Request{UserConfigInfo: info, Cfg: cfg}
	return request
}

func (self *Request) Run() *goquery.Document {
	return self.http_request()
}

func (self *Request) http_request() *goquery.Document {
	var tr *http.Transport
	if self.Cfg.Proxy.IsUse == 1 {
		//先临时固定代理地址
		proxy := func(_ *http.Request) (*url.URL, error) {
			return url.Parse(self.Cfg.Proxy.Link)
		}

		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           proxy,
			//DisableCompression: true,
		}
	} else {
		tr = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
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

	b, _ := json.Marshal(self.UserConfigInfo.Param)
	body := bytes.NewBuffer(b)
	if len(self.UserConfigInfo.Param) < 1 {
		//resp, _ = c.Get(self.UserConfig.Url)
		req, _ = http.NewRequest("GET", self.UserConfigInfo.Url, nil)
	} else {
		//resp, _ = c.Post(self.UserConfig.Url, "application/x-www-form-urlencoded", body)
		req, _ = http.NewRequest("POST", self.UserConfigInfo.Url, body)
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
	resp, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	res, _ := goquery.NewDocumentFromResponse(resp)
	//fmt.Println(res)
	//os.Exit(-1)
	return res
}
