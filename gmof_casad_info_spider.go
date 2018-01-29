package main

import (
	"log"

	"github.com/celrenheit/spider"
)

type GMOF_Person_Casad_HTMLSpider struct {
	title string `json:"title"`
	url   string `json:"url"`
	desc  string `json:"desc"`
}

func NewGMOF_Person_Casad_HTMLSpider() *GMOF_Person_Casad_HTMLSpider {
	return &GMOF_Person_Casad_HTMLSpider{
		desc: "",
	}
}

func (w *GMOF_Person_Casad_HTMLSpider) Setup(ctx *spider.Context) (*spider.Context, error) {
	return spider.NewHTTPContext("GET", w.url, nil)
}

func (w *GMOF_Person_Casad_HTMLSpider) Spin(ctx *spider.Context) error {
	if _, err := ctx.DoRequest(); err != nil {
		return err
	}

	html, err := ctx.HTMLParser()
	if err != nil {
		return err
	}

	var person GMOF_Person_Casad
	person.Sourceurl = w.url

	//<title></title>
	title := html.Find("title").Eq(0).Text()

	name := html.Find(".contentBar .title h1").Eq(0).Text()
	person.Name = name

	//Style A
	//误操作丢失

	log.Println("GMOF_Casad_Info_HTMLSpider Find title:[%s]", title)
	log.Println("GMOF_Casad_Info_HTMLSpider Find Name:[%s]", person.Name)
	log.Println("GMOF_Casad_Info_HTMLSpider Find Birthday:[%s]", person.Birthday)
	log.Println("GMOF_Casad_Info_HTMLSpider Find Home:[%s]", person.Home)
	log.Println("GMOF_Casad_Info_HTMLSpider Find Resume:[%s]", person.Resume)
	log.Println("GMOF_Casad_Info_HTMLSpider Find() Finish--[%s]", w.url)

	if person.Birthday == "" {
		//panic("Birthday is null......")
	}

	go Save_GMOF_Person_Casad(person)
	return err
}
