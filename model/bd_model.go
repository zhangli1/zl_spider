package model

import (
	"zl_spider/config"
)

type BdModel struct {
	Coding config.UserConfigInfo
	Cfg    config.Config
}

func NewBdModel(coding config.UserConfigInfo, cfg config.Config) *BdModel {
	bd_model := &BdModel{Coding: coding, Cfg: cfg}
	return bd_model
}

func (self *BdModel) Run() interface{} {
	return self.Destruct("test")
}

func (self *BdModel) Destruct(param interface{}) interface{} {
	ret_data := make(map[string]interface{})
	ret_data["bd"] = param
	return ret_data

}
