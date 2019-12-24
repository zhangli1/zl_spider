/*
 *   入口文件
 */

package main

import (
	"zl_spider/config"
	"zl_spider/controller"

	"gopkg.in/gcfg.v1"
)

func main() {
	var server_cfg config.Config
	gcfg.ReadFileInto(&server_cfg, "conf/base.cfg")

	dir := "/test"
	spider := controller.NewSpider(dir, server_cfg)
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
