package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/celrenheit/spider"
	"github.com/celrenheit/spider/schedule"
)

var (
	// Ensure WikipediaHTMLSpider implements spider.Spider interface
	_ spider.Spider = (*GMOF_CaiPiao_Month_HTMLSpider)(nil)
)

func NewTTankTask(taskname string) {

	if taskname == "GMOF_CaiPiao_Default" {
		url := "http://zhs.mof.gov.cn/zhuantilanmu/caipiaoguanli/201712/t20171221_2786493.html"
		url = "http://zhs.mof.gov.cn/zhuantilanmu/caipiaoguanli//201204/t20120417_643575.html"

		caipiaoSpider := init_GMOF_CaiPiao_Month_HTMLSpider(url)
		ctx, _ := caipiaoSpider.Setup(nil)
		caipiaoSpider.Spin(ctx)
	} else if taskname == "GMOF_CaiPiao_CSV" {
		data := read_csv("./data/Data_GMOF_CaiPiao_List2.csv", ",")
		if data != nil {
			for i := range data {
				var url = "http://zhs.mof.gov.cn/zhuantilanmu/caipiaoguanli/"
				row := strings.Replace(data[i], "./", "", -1)
				url = url + row
				fmt.Printf("========%s \n", url)

				caipiaoSpider := init_GMOF_CaiPiao_Month_HTMLSpider(url)
				ctx, _ := caipiaoSpider.Setup(nil)
				caipiaoSpider.Spin(ctx)

				spider.Add(schedule.Every(5*time.Second), caipiaoSpider)
				spider.Start()
			}
		}

	} else if taskname == "GMOF_CaiPiao_List" {
		data := make([]string, 15)

		for i := range data {
			url := "http://zhs.mof.gov.cn/zhuantilanmu/caipiaoguanli/"
			if i > 0 {
				url = url + "index_" + fmt.Sprintf("%d", i) + ".html"
			}
			log.Printf(">>>>>%s\n", url)
			caipiaoSpider := init_GMOF_CaiPiao_List_HTMLSpider(url)
			ctx, _ := caipiaoSpider.Setup(nil)
			caipiaoSpider.Spin(ctx)

			spider.Add(schedule.Every(5*time.Second), caipiaoSpider)
			spider.Start()
		}

		//<-time.After(8 * time.Second)
	} else {
		fmt.Println("----------- can not match this taskname:" + taskname)
	}

}
