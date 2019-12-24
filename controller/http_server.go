/*
 *  提供给外部的http接口
 *  2017-10-28
 */

package controller

import ()

type HttpServer struct {
}

func NewHttpServer() *HttpServer {
     http_server := &HttpServer{}
     return http_server
}

func (self *HttpServer) Run() interface{} {
    return "http_server"
}

