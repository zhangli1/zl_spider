package controller

import ()

type Download struct {

}

func NewDownload() *Download {
    download := &Download{}
    return download
}


func (self *Download) Run() interface{} {
    return "download"
}
