package config

type Config struct {
    Base struct {
        Url string
        Timeout int64
    }
}


type UserConfigInfo struct {
    Url string
    Timeout int64
    Param string
}



func GetUserConfigInfo() []UserConfigInfo {
    var user_config_info_list []UserConfigInfo
    user_config_info_list = []UserConfigInfo{
        UserConfigInfo{Url:"https://www.baidu.com", Timeout:10, Param:""},
        UserConfigInfo{Url:"http://www.zhangli.me", Timeout:15, Param:""},
    }
    return user_config_info_list
}


