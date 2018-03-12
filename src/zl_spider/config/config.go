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
    Action string
    Coding string
}



func GetUserConfigInfo() []UserConfigInfo {
    var user_config_info_list []UserConfigInfo
    user_config_info_list = []UserConfigInfo{
        UserConfigInfo{Url:"https://www.baidu.com", Timeout:10, Param:"", Action:"bd"},
        //UserConfigInfo{Url:"http://www.zhangli.me", Timeout:15, Param:"", Action:"zl"},
        //UserConfigInfo{Url:"http://newhouse.wuhan.fang.com/house/s/jiangxia2/a77-b25000%2C10000-b82-c412/", Timeout:25, Param:"", Action:"fang", Coding:"GBK"},
    }
    return user_config_info_list
}


