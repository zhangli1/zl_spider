/*
 *   入口文件
 */

 package main

 import (
     "fmt"
     "zl_spider/controller"
 )

 func main(){
     dir := "/tmp/test"
     config := "config_test"
     spider := controller.NewSpider(dir, config)
    fmt.Println(spider.Run())
 }
