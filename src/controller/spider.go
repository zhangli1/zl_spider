/*
 * zl_spider主入口文件
 */


package spider

import (
)

type Spider struct {
    ExeDir string
    Cfg    config.Config
}

func NewSpider(exeDir string, cfg config.Config) (spider *Spider) {
    spider = &Spider{ExeDir : exeDir, Cfg : cfg}
    return spider
}
