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

func init_GMOF_CDC_Epide_List_HTMLSpider(url string) *GMOF_CDC_Epide_List_HTMLSpider {
	spider := NewGMOF_CDC_Epide_List_HTMLSpider()
	spider.url = url
	return spider
}

func init_GMOF_CDC_Epide_Month_HTMLSpider(url string) *GMOF_CDC_Epide_Month_HTMLSpider {
	spider := NewGMOF_CDC_Epide_Month_HTMLSpider()
	spider.url = url
	return spider
}

func init_GMOF_Person_Region_HTMLSpider(url string) *GMOF_Person_Region_HTMLSpider {
	spider := NewGMOF_Person_Region_HTMLSpider()
	spider.url = url
	return spider
}

func init_GMOF_Casad_List_HTMLSpider(url string) *GMOF_Casad_List_HTMLSpider {
	spider := NewGMOF_Casad_List_HTMLSpider()
	spider.url = url
	return spider
}

func init_GMOF_Person_Casad_HTMLSpider(url string) *GMOF_Person_Casad_HTMLSpider {
	spider := NewGMOF_Person_Casad_HTMLSpider()
	spider.url = url
	return spider
}
