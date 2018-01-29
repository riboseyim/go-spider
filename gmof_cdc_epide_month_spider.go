package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/celrenheit/spider"
)

type GMOF_CDC_Epide_Month_HTMLSpider struct {
	title string `json:"title"`
	url   string `json:"url"`
	desc  string `json:"desc"`
}

func NewGMOF_CDC_Epide_Month_HTMLSpider() *GMOF_CDC_Epide_Month_HTMLSpider {
	return &GMOF_CDC_Epide_Month_HTMLSpider{
		desc: "",
	}
}

func (w *GMOF_CDC_Epide_Month_HTMLSpider) Setup(ctx *spider.Context) (*spider.Context, error) {
	log.Println("---url----%s", w.url)
	return spider.NewHTTPContext("GET", w.url, nil)
}

func (w *GMOF_CDC_Epide_Month_HTMLSpider) Spin(ctx *spider.Context) error {
	if _, err := ctx.DoRequest(); err != nil {
		return err
	}

	html, err := ctx.HTMLParser()

	if err != nil {
		return err
	}

	cdc_epide := NewGMOF_CDC_Epide_Month()

	//<title></title>
	cdc_epide.Title = html.Find("title").Eq(0).Text()
	//	cdc_epide.Title = Convert2String(cdc_epide.Title, GB18030)
	log.Println("---title----%s", cdc_epide.Title)
	test, _ := html.Html()
	log.Println("---html----%s", test)

	// 问题：ajax or js encode

	html.Find("#xw_box > strong > font > table > tbody").Each(func(i int, s *goquery.Selection) {
		tr := s.Find("tr").Text()
		log.Println("---tr----%s", tr)
		cdc_epide.Name = html.Find("td p").Eq(0).Text()
		cdc_epide.Incidence = html.Find("td p").Eq(1).Text()
		cdc_epide.Death = html.Find("td p").Eq(2).Text()

	})

	log.Println("GMOF_CDC_Epide_Month_HTMLSpider match cdc_epide:%s,%s,%s", cdc_epide.Name, cdc_epide.Incidence, cdc_epide.Death)
	log.Println("---------GMOF_CDC_Epide_Month_HTMLSpider Find() Finish----------")

	saveData_GMOF_CDC_Epide_Month(cdc_epide)
	return err
}
