package controller

import (
	"zl_spider/config"
)

type Rule struct {
}

func NewRule() *Rule {
	rule := &Rule{}
	return rule
}

func (self *Rule) Run(cfg config.Config) []config.UserConfigInfo {
    return config.GetUserConfigInfo()
}
