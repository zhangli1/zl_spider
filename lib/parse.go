package lib

import (
	//"fmt"
	//"reflect"
	"zl_spider/config"
	//"zl_spider/model"

	"github.com/PuerkitoBio/goquery"
	//"github.com/Tang-RoseChild/mahonia"
)

type Parse struct {
	Resp   *goquery.Document
	Coding config.UserConfigInfo
	Cfg    config.Config
}

func NewParse(resp *goquery.Document, coding config.UserConfigInfo, cfg config.Config) *Parse {
	parse := &Parse{Resp: resp, Coding: coding, Cfg: cfg}
	return parse
}

/*func (self *Parse) Run() interface{} {
	model_list := make(map[string]model.Model, 0)
	model_list["boss"] = model.NewBossModel(self.Resp, self.Cfg)
	//model_list["bd"] = model.NewBdModel(self.Resp, self.Cfg)

	if model_list[self.Coding.ModelPrefix] != nil {
		return model_list[self.Coding.ModelPrefix].Run()
	}

	return ""
}*/
