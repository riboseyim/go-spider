package main

import (
	"log"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/celrenheit/spider"
)

type GMOF_Casad_List_HTMLSpider struct {
	title string `json:"title"`
	url   string `json:"url"`
	desc  string `json:"desc"`
}

func NewGMOF_Casad_List_HTMLSpider() *GMOF_Casad_List_HTMLSpider {
	return &GMOF_Casad_List_HTMLSpider{
		desc: "",
	}
}

func (w *GMOF_Casad_List_HTMLSpider) Setup(ctx *spider.Context) (*spider.Context, error) {
	return spider.NewHTTPContext("GET", w.url, nil)
}

func (w *GMOF_Casad_List_HTMLSpider) Spin(ctx *spider.Context) error {
	if _, err := ctx.DoRequest(); err != nil {
		return err
	}

	html, err := ctx.HTMLParser()
	if err != nil {
		return err
	}

	GMOF_Casad_Array := make([]GMOF_Person_Casad, 1000)
	GMOF_Casad_Array_Index := 0
	html.Find("#allNameBar").Each(func(i int, s *goquery.Selection) {
		group := s.Find("dt").Text()
		s.Find("dd span").Each(func(i int, s *goquery.Selection) {
			name := s.Find("a").Eq(0).Text()
			href, _ := s.Find("a").Eq(0).Attr("href")

			var person GMOF_Person_Casad
			person.Title = "中国科学院-" + group + "-" + name
			person.Name = name
			person.Sourceurl = href

			if person.Sourceurl != "" && strings.Index(person.Sourceurl, "www.casad.cas.cn") > 0 {
				GMOF_Casad_Array[GMOF_Casad_Array_Index] = person
				GMOF_Casad_Array_Index++
				if GMOF_Casad_Array_Index%50 == 0 {
					time.Sleep(5 * time.Second)
				}
				go Handle_GMOF_Casad_Info_Task(person.Sourceurl)

				//panic("----Test Return-------")

			}
		})

	})
	//casad_list := NewGMOF_Casad_List()

	//<title></title>
	title := html.Find("title").Eq(0).Text()

	log.Println("GMOF_Casad_List_HTMLSpider match title:%s", title)
	log.Println("GMOF_Casad_List_HTMLSpider Find() Finish----------")
	return err
}
