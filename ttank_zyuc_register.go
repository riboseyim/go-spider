package main

func init_ZYUC_Task_Info_HTMLSpider(url string) *ZYUC_Task_Info_HTMLSpider {
	spider := NewZYUC_Task_Info_HTMLSpider()
	spider.url = url
	return spider
}
