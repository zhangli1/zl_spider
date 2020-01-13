/*
 *   入口文件
 */

package main

import (
	"fmt"
	"runtime/debug"
	"time"
	"zl_spider/config"
	"zl_spider/controller"

	"gopkg.in/gcfg.v1"
)

func main() {
	var server_cfg config.Config
	gcfg.ReadFileInto(&server_cfg, "conf/base.cfg")
	dir := "/test"
	spider := controller.NewSpider(dir, server_cfg)

	defer func() {
		if re := recover(); re != nil {
			debug.PrintStack()
			fmt.Println(string(debug.Stack()))
			time.Sleep(time.Duration(3) * time.Second)

		}
	}()
	spider.Run()

	/*for _, v := range ret.(map[int]interface{}) {
		var list []lib.JobData
		err := json.Unmarshal([]byte(v.(string)), &list)
		if err != nil {
			fmt.Println("parse json fail")
		}
		fmt.Println(list[0].JobTitle)
	}*/
}
