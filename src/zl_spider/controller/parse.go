package controller

import (
    "github.com/PuerkitoBio/goquery"
)

type Parse struct {
    Resp *goquery.Document
}

func NewParse(resp *goquery.Document) *Parse {
    parse := &Parse{Resp : resp}
    return parse
}

func (self *Parse) Run() interface{} {
    var href string
    list := make([]string, 0)
    self.Resp.Find("div").Each(func(i int, s *goquery.Selection) {
        href, _ = s.Attr("class")
        list = append(list, href)
    })
    return list
}

/*func (self *Parse) parse_str() string {

}*/
