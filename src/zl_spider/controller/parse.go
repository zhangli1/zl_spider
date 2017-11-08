package controller

import (
    "github.com/PuerkitoBio/goquery"
)

type Parse struct {
    Content string
}

func NewParse(content string) *Parse {
    parse := &Parse{Content : content}
    return parse
}

func (self *Parse) Run() string {
    return self.Content
}

func (self *Parse) parse_str() string {

}
