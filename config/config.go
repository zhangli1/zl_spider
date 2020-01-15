package config

type Config struct {
	Base struct {
		Url     string
		Timeout int64
	}
	Redis struct {
		Host   string
		Port   string
		Passwd string
		Select int
	}
	Elasticsearch struct {
		Url string
	}
	Frequency struct {
		City int
	}
	Proxy struct {
		Links string
		IsUse int
	}
}

type UserConfigInfo struct {
	Url         string                 //链接
	Timeout     int64                  //超时时间
	Param       map[string]interface{} //参数
	Action      string                 //方法类型
	Coding      string
	ModelPrefix string //模型实例化前缀
	Switch      bool   //开关， 是否打开抓取
}

func GetUserConfigInfo() []UserConfigInfo {
	var user_config_info_list []UserConfigInfo
	user_config_info_list = []UserConfigInfo{
		UserConfigInfo{
			Url:         "https://www.zhipin.com/mobile/jobs.json",
			Timeout:     15,
			ModelPrefix: "boss",
			Action:      "",
			Switch:      true,
		},
		UserConfigInfo{
			Url:         "http://www.feixiaohao.com/list_2.html",
			Timeout:     15,
			ModelPrefix: "feixiaohao",
			Action:      "",
			Switch:      false,
		},
		UserConfigInfo{
			Url:         "https://www.zhihu.com/search?type=content&q=%E6%B7%B1%E5%9C%B3",
			Timeout:     15,
			ModelPrefix: "zhihu",
			Action:      "",
			Switch:      false,
		},
		UserConfigInfo{
			Url:         "https://www.v2ex.com",
			Timeout:     25,
			ModelPrefix: "v2ex",
			Action:      "fang",
			Switch:      false,
		},
		UserConfigInfo{
			Url:         "https://www.freeip.top/api/proxy_ips",
			Timeout:     15,
			ModelPrefix: "freeip",
			Action:      "freeip",
			Switch:      true,
		},
	}
	return user_config_info_list
}
