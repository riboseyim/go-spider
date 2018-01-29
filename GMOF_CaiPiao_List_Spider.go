package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/celrenheit/spider"
	uuid "github.com/satori/go.uuid"
)

type GMOF_CaiPiao_List_HTMLSpider struct {
	title string `json:"title"`
	url   string `json:"url"`
	desc  string `json:"desc"`
}

func NewGMOF_CaiPiao_List_HTMLSpider() *GMOF_CaiPiao_List_HTMLSpider {
	return &GMOF_CaiPiao_List_HTMLSpider{
		desc: "",
	}
}

func (w *GMOF_CaiPiao_List_HTMLSpider) Setup(ctx *spider.Context) (*spider.Context, error) {
	return spider.NewHTTPContext("GET", w.url, nil)
}

func (w *GMOF_CaiPiao_List_HTMLSpider) Spin(ctx *spider.Context) error {
	if _, err := ctx.DoRequest(); err != nil {
		return err
	}

	html, err := ctx.HTMLParser()
	if err != nil {
		return err
	}

	list := NewGMOF_CaiPiao_List()

	//<title></title>
	list.Title = html.Find("title").Eq(0).Text()
	list.Title = Convert2String(list.Title, GB18030)
	//class="TRS_Editor"
	html.Find(".RdZt_CPiao_R_mid_Zi").Each(func(i int, s *goquery.Selection) {
		title := s.Find("span a").Text()
		href, _ := s.Find("span a").Eq(0).Attr("href")

		title = Convert2String(title, GB18030)
		if title != "" && strings.Index(title, "全国彩票销售情况") > 0 {
			list.Title = title
			list.Url = href

			event := new(GMOF_CaiPiao_List)
			event.AccId = uuid.NewV4().String()
			event.Title = list.Title
			event.Type = "Add_GMOF_CaiPiao_List"
			event.Url = list.Url

			//=======kafka
			//kafka := newKafkaSyncProducer()
			//sendMsg(kafka, topic_ttank_gmof_caipiao_list, event)

			saveData_GMOF_CaiPiao_List(event)
		}
	})

	log.Printf("---------GMOF_CaiPiao_List_HTMLSpider Find() Finish----------")

	return err
}
