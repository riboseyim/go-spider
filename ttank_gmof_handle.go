package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/celrenheit/spider"
)

var (
	// Ensure  implements spider.Spider interface
	_ spider.Spider = (*GMOF_CaiPiao_Month_HTMLSpider)(nil)
	_ spider.Spider = (*GMOF_CaiPiao_List_HTMLSpider)(nil)
	_ spider.Spider = (*GMOF_CDC_Epide_Month_HTMLSpider)(nil)
	_ spider.Spider = (*GMOF_CDC_Epide_List_HTMLSpider)(nil)
	_ spider.Spider = (*GMOF_Person_Region_HTMLSpider)(nil)
	_ spider.Spider = (*GMOF_Person_Casad_HTMLSpider)(nil)
)

func Handle_GMOF_CDC_Epide_List_Task() {
	url := "http://www.casad.cas.cn/chnl/371/index.html"
	if url != "" {
		myspider := init_GMOF_CDC_Epide_List_HTMLSpider(url)
		ctx, _ := myspider.Setup(nil)
		myspider.Spin(ctx)
	}

	<-time.After(300 * time.Second)
}

func Handle_GMOF_CDC_Epide_Month_Test() {
	url := "http://www.nhfpc.gov.cn/jkj/s3578/201801/178264d9c8ab4d439a34284a4c84fee7.shtml"
	Handle_GMOF_CDC_Epide_Month_Task(url)
}

func Handle_GMOF_CDC_Epide_Month_Task(url string) {
	if url != "" {
		myspider := init_GMOF_CDC_Epide_Month_HTMLSpider(url)
		ctx, _ := myspider.Setup(nil)
		myspider.Spin(ctx)
	}
}

func Handle_GMOF_Casad_ReQuery_Task() {
	persons := query_GMOF_person_Casad_NoHome()
	for key := range persons {
		url := persons[key]
		Delete_GMOF_Person_Casad(url)
		if url != "" {
			myspider := init_GMOF_Casad_List_HTMLSpider(url)
			ctx, _ := myspider.Setup(nil)
			myspider.Spin(ctx)
		}
	}

	<-time.After(300 * time.Second)
}
func Handle_GMOF_Casad_List_Task() {
	url := "http://www.casad.cas.cn/chnl/371/index.html"
	if url != "" {
		myspider := init_GMOF_Casad_List_HTMLSpider(url)
		ctx, _ := myspider.Setup(nil)
		myspider.Spin(ctx)
	}

	<-time.After(300 * time.Second)
}

func Handle_GMOF_Casad_Info_Test() {
	url := "http://www.casad.cas.cn:80/aca/371/dxb-201711-t20171129_4625076.html"
	url = "http://www.casad.cas.cn:80/aca/371/xxjskxb-200906-t20090624_1807804.html"
	url = "http://www.casad.cas.cn:80/aca/371/xxjskxb-200906-t20090624_1807718.html"
	url = "http://www.casad.cas.cn/aca/371/xxjskxb-201112-t20111213_3412392.html"
	Handle_GMOF_Casad_Info_Task(url)
}

func Handle_GMOF_Casad_Info_Task(url string) {
	if url != "" {
		myspider := init_GMOF_Person_Casad_HTMLSpider(url)
		ctx, _ := myspider.Setup(nil)
		myspider.Spin(ctx)
	}
}

func Handle_GMOF_Person_Region_BatchTask() {
	//4915-personPage5721-10000--11598
	//10001-11598
	// 2017-01-07,似乎有反爬虫设置，可以设置500个一批次并发
	//4900
	for i := 11400; i < 12000; i++ {
		url := "http://ldzl.people.com.cn/dfzlk/front/personPage"
		if i > 0 {
			url = url + fmt.Sprintf("%d", i) + ".htm"
			log.Printf(">>>>>%s\n", url)

			if i%200 == 0 {
				time.Sleep(10 * time.Second)
			}
			go Handle_GMOF_Person_Region_Task(url)
		}
	}
	log.Printf(">>>>>Handle_GMOF_Person_Region_BatchTask() finish \n")
	<-time.After(300 * time.Second)
}

func Handle_GMOF_Person_Region_Test() {
	url := "http://ldzl.people.com.cn/dfzlk/front/personPage11479.htm"
	Handle_GMOF_Person_Region_Task(url)
}

func Handle_GMOF_Person_Region_Task(url string) {
	if url != "" {
		myspider := init_GMOF_Person_Region_HTMLSpider(url)
		ctx, _ := myspider.Setup(nil)
		myspider.Spin(ctx)
	}
}

func Handle_GMOF_CaiPiao_Month_BatchTask() {
	data := read_csv_caipiao("./data/Data_GMOF_CaiPiao_List.csv", ",")
	if data != nil {
		for i := range data {
			row := strings.Replace(data[i], "../", "", -1)
			row = strings.Replace(row, "./", "", -1)
			row = strings.Replace(row, "..", "", -1)

			url := ""
			if strings.Contains(row, "http") {
				url = row
			} else if strings.Contains(row, "zhengwuxinxi") {
				url = "http://zhs.mof.gov.cn/"
				url = url + row
			} else {
				url = "http://zhs.mof.gov.cn/zhuantilanmu/caipiaoguanli/"
				url = url + row
			}
			fmt.Printf("%d=====%s \n", i, url)
			go Handle_GMOF_CaiPiao_Month_Task(url)
		}
		<-time.After(60 * time.Second)
	}
}

func Handle_GMOF_CaiPiao_Month_Test() {
	url := "http://zhs.mof.gov.cn/zhuantilanmu/caipiaoguanli/201712/t20171221_2786493.html"
	//url = "http://zhs.mof.gov.cn/zhuantilanmu/caipiaoguanli//201204/t20120417_643575.html"
	Handle_GMOF_CaiPiao_Month_Task(url)
}

func Handle_GMOF_CaiPiao_Month_Task(url string) {
	if url != "" {
		myspider := init_GMOF_CaiPiao_Month_HTMLSpider(url)
		ctx, _ := myspider.Setup(nil)
		myspider.Spin(ctx)
	}
}

func Handle_GMOF_CaiPiao_List_BatchTask() {
	data := make([]string, 15)
	for i := range data {
		url := "http://zhs.mof.gov.cn/zhuantilanmu/caipiaoguanli/"
		if i > 0 {
			url = url + "index_" + fmt.Sprintf("%d", i) + ".html"
		}
		log.Printf(">>>>> %s \n", url)
		go Handle_GMOF_CaiPiao_List_Task(url)
	}
	<-time.After(10 * time.Second)
}

func Handle_GMOF_CaiPiao_List_Task(url string) {
	if url != "" {
		myspider := init_GMOF_CaiPiao_List_HTMLSpider(url)
		ctx, _ := myspider.Setup(nil)
		myspider.Spin(ctx)
	}
}
