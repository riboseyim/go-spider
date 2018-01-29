package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/celrenheit/spider"
)

type GMOF_Person_Region_HTMLSpider struct {
	title string `json:"title"`
	url   string `json:"url"`
	desc  string `json:"desc"`
}

func NewGMOF_Person_Region_HTMLSpider() *GMOF_Person_Region_HTMLSpider {
	return &GMOF_Person_Region_HTMLSpider{
		desc: "",
	}
}

func (w *GMOF_Person_Region_HTMLSpider) Setup(ctx *spider.Context) (*spider.Context, error) {
	return spider.NewHTTPContext("GET", w.url, nil)
}

func (w *GMOF_Person_Region_HTMLSpider) Spin(ctx *spider.Context) error {
	if _, err := ctx.DoRequest(); err != nil {
		return err
	}

	html, err := ctx.HTMLParser()
	if err != nil {
		return err
	}

	person := NewGMOF_Person_Region()

	//<title></title>
	title := html.Find("title").Eq(0).Text()
	if title != "" {
		title = Convert2String(title, GB18030)
		names := strings.Split(title, " ")
		if len(names) > 0 {
			name := names[0]
			log.Println("====Name:%s", name)
			person.Name = name
		}
	}

	if person.Name == "404" {
		return nil
	}

	position := html.Find(".clearfix dd span").Text()
	log.Printf("----position---- %s", Convert2String(position, UTF8))
	if position != "" {
		person.Title = Convert2String(position, GB18030)
	}

	//class="p2j_text"
	recordsIndex := 0
	records := make([]string, 100)
	html.Find(".p2j_text p").Each(func(i int, s *goquery.Selection) {
		record := s.Text()
		if record != "" {
			log.Printf("----record----%d %s", recordsIndex, Convert2String(record, UTF8))

			person.Resume = person.Resume + Convert2String(record, GB18030)

			//杨永英，女，布依族，贵州荔波人，1964年4月生，省委党校研究生学历，1985年11月参加工作，1984年11月加入中国共产党。

			if recordsIndex == 0 {
				record = Convert2String(record, GB18030)
				person.Summary = record
				commons := strings.Split(record, "，")
				for i := range commons {
					log.Printf("----commons----%d %s", i, Convert2String(commons[i], UTF8))
					record = Convert2String(commons[i], GB18030)
					if commons[i] == "男" || commons[i] == "女" {
						person.Sex = commons[i]
					}
					if strings.Index(commons[i], "族") > 0 {
						person.Ethnic = commons[i]
					}
					if strings.Index(commons[i], "学历") > 0 || strings.Index(commons[i], "学位") > 0 || strings.Index(commons[i], "毕业") > 0 {
						person.Education = person.Education + "," + commons[i]
					}
					if strings.Index(commons[i], "人") > 0 {
						person.Home = commons[i]
					}
					if strings.Index(commons[i], "生") > 0 {
						person.Birthday = commons[i]
					}
					if strings.Index(commons[i], "参加工作") > 0 {
						person.Workday = commons[i]
					}
					if strings.Index(commons[i], "加入") > 0 {
						person.Partyday = commons[i]
					}
				}
			}

			records[recordsIndex] = record
			recordsIndex++
		}
	})
	//person.Records = records

	//id="appendix"
	html.Find("#userimg").Each(func(i int, s *goquery.Selection) {
		//href, _ := s.Find("img").Eq(0).Attr("src") //附件
		//log.Printf("=====appendix %d:======== \n %s \n ", i, href)
		//person.Resume = person.Resume + "\n" + Convert2String(href, GB18030)
	})
	person.Sourceurl = w.url
	log.Printf("---------GMOF_Person_HTMLSpider Find() Finish----------")

	fmt.Printf("-------finish %+v -------", person)
	saveData_GMOF_Person_Region(person)
	return err
}
