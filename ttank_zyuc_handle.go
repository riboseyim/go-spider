package main

import (
	"github.com/celrenheit/spider"
)

var (
	// Ensure  implements spider.Spider interface
	_ spider.Spider = (*ZYUC_Task_Info_HTMLSpider)(nil)
)

//疫情
//http://www.nhfpc.gov.cn/jkj/s3578/new_list_2.shtml
//电话用户分省情况
//http://www.miit.gov.cn/newweb/n1146312/n1146904/n1648372/c5966257/content.html

func Handle_ZYUC_Task_Info_Collect() {
	clear_csv("./zyuc/Task_List_New.cfg")
	tasks := read_csv_normal("./zyuc/tasklist.cfg")
	for i := range tasks {
		taskid := tasks[i]
		url := "http://192.168.6.67:2014/tasksys/TaskDetail.jsp?taskID=" + taskid
		Handle_ZYUC_Task_Info_Task(url)
	}

}

func Handle_ZYUC_Task_Info_Compare(url string) {
	if url != "" {
		myspider := init_ZYUC_Task_Info_HTMLSpider(url)
		ctx, _ := myspider.Setup(nil)
		myspider.Spin(ctx)
	}
}

func Handle_ZYUC_Task_Info_Task(url string) {
	if url != "" {
		myspider := init_ZYUC_Task_Info_HTMLSpider(url)
		ctx, _ := myspider.Setup(nil)
		myspider.Spin(ctx)
	}
}
