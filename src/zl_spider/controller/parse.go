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
    list := make([]string, 0)
    item := self.Resp.Find("div[class=\"List-item\"]")
    item.Each(func(i int, s *goquery.Selection) {
        text := s.Eq(i).Find("span[class=\"RichText CopyrightRichText-richText\"]").Text()
        list = append(list, text)
    })
    return list
}

/*func (self *Parse) parse_str() string {

}*/
