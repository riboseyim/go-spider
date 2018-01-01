package main

func init_GMOF_CaiPiao_List_HTMLSpider(url string) *GMOF_CaiPiao_List_HTMLSpider {
	spider := NewGMOF_CaiPiao_List_HTMLSpider()
	spider.url = url
	return spider
}

func init_GMOF_CaiPiao_Month_HTMLSpider(url string) *GMOF_CaiPiao_Month_HTMLSpider {
	spider := NewGMOF_CaiPiao_Month_HTMLSpider()
	spider.url = url
	return spider
}
