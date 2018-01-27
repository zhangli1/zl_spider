package controller

import (
	"zl_spider/config"
)

type Rule struct {
}

//需要用户操心的返回结构
type UserConfig struct {
	Url     string
	TimeOut int64
	Param   map[string]interface{}
}

func NewRule() *Rule {
	rule := &Rule{}
	return rule
}

/*func (self *Rule) getRule() interface{} {

}*/

func (self *Rule) Run(cfg config.Config) UserConfig {
	var userConfig UserConfig

	userConfig.Url = cfg.Base.Url 
	userConfig.TimeOut = cfg.Base.Timeout
	return userConfig
}
