/*
 *   入口文件
 */

package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"zl_spider/config"
	"zl_spider/controller"

	l4g "code.google.com/p/log4go"
	"gopkg.in/gcfg.v1"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	pwd, _ := os.Getwd()
	executeDir := *flag.String("d", pwd, "execute directory") + "/"

	configFilePath := flag.String("c", "conf/base.cfg", "config file")
	flag.Parse()
	fmt.Println("Current execute directory:", executeDir)
	globalLogger := l4g.NewDefaultLogger(l4g.DEBUG)
	globalLoggerBoss := l4g.NewDefaultLogger(l4g.DEBUG)
	globalLoggerBoss.LoadConfiguration(executeDir + "conf/l4g_boss.xml")
	globalLogger.LoadConfiguration(executeDir + "conf/l4g_free.xml")

	configFile := executeDir + *configFilePath
	/*c, err := ReadConfigFile(config)
	if err != nil {
		fmt.Println("Get config file:[" + config + "] failed, exiting.")
		os.Exit(-1)
	}*/
	//启动信号监听
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	var server_cfg config.Config
	gcfg.ReadFileInto(&server_cfg, configFile)
	dir := executeDir

	loggerConfigMap := make(map[string]*l4g.Logger, 0)
	loggerConfigMap["boss"] = &globalLoggerBoss
	loggerConfigMap["freeip"] = &globalLogger
	spider := controller.NewSpider(dir, server_cfg, loggerConfigMap, sigs)

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
