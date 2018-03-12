package controller

import (
//	"fmt"
//	"os"
    "github.com/PuerkitoBio/goquery"
    "github.com/Tang-RoseChild/mahonia"
)

type Parse struct {
    Resp *goquery.Document
    Coding string
}

func NewParse(resp *goquery.Document, coding string) *Parse {
    parse := &Parse{Resp : resp, Coding : coding}
    return parse
}

func (self *Parse) Run() interface{} {
    item, _ := self.Resp.Find("body").Html()
    if self.Coding == "GBK" {
        enc:=mahonia.NewDecoder("GBK")
        return enc.ConvertString(item)
    }
    return item
}
