/*
 *   入口文件
 */

 package main

 import (
     "fmt"
     "zl_spider/controller"
     "gopkg.in/gcfg.v1"
	 "zl_spider/config"
 )

func main(){
	var server_cfg config.Config
    gcfg.ReadFileInto(&server_cfg, "conf/base.cfg")

	dir := "/test"
    spider := controller.NewSpider(dir, server_cfg)
    fmt.Println(spider.Run())
 }
