package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/celrenheit/spider"
)

type GMOF_CaiPiao_Month_HTMLSpider struct {
	title string `json:"title"`
	url   string `json:"url"`
	desc  string `json:"desc"`
}

func NewGMOF_CaiPiao_Month_HTMLSpider() *GMOF_CaiPiao_Month_HTMLSpider {
	return &GMOF_CaiPiao_Month_HTMLSpider{
		desc: "",
	}
}

func (w *GMOF_CaiPiao_Month_HTMLSpider) Setup(ctx *spider.Context) (*spider.Context, error) {
	return spider.NewHTTPContext("GET", w.url, nil)
}

func (w *GMOF_CaiPiao_Month_HTMLSpider) Spin(ctx *spider.Context) error {
	if _, err := ctx.DoRequest(); err != nil {
		return err
	}

	html, err := ctx.HTMLParser()
	if err != nil {
		return err
	}

	caipiao := NewGMOF_CaiPiao_Month()

	//<title></title>
	caipiao.Title = html.Find("title").Eq(0).Text()
	caipiao.Title = Convert2String(caipiao.Title, GB18030)
	//class="TRS_Editor"
	html.Find(".TRS_Editor").Each(func(i int, s *goquery.Selection) {
		content := s.Find("p").Text()
		caipiao.Content = content

		if content != "" {
			content = Convert2String(content, GB18030)
			rows := strings.Split(content, "。")

			for _, value := range rows {
				//fmt.Printf("======arr[%d]=\n [%s] \n", index, value)
				if strings.Index(value, "全国彩票") > 0 {
					reg := regexp.MustCompile(`全国共销售彩票([\d]+.[\d]+)\S+`)
					result := reg.FindStringSubmatch(value)
					if len(result) > 0 {
						caipiao.Total = result[1]
					}
				}
			}
		}

	})
	//id="appendix"
	html.Find("#appendix").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Find("a").Eq(0).Attr("href") //附件
		//log.Printf("=====appendix %d:======== \n %s \n ", i, href)
		caipiao.Attachid = href
	})

	if caipiao.Total == "0" {
		log.Printf("---------GMOF_CaiPiao_HTMLSpider 201204之前的样式----------")
		//class="TRS_Editor"
		content := html.Find(".Custom_UnionStyle").Text()

		if content != "" {
			content = strings.Replace(content, "<P>", "", -1)
			content = strings.Replace(content, "<SPAN>", "", -1)
			content = strings.Replace(content, "<STRONG>", "", -1)
			content = strings.Replace(content, "</P>", "", -1)
			content = strings.Replace(content, "</SPAN>", "", -1)
			content = strings.Replace(content, "</STRONG>", "", -1)

			content = Convert2String(content, GB18030)
			rows := strings.Split(content, "。")

			for _, value := range rows {
				//fmt.Printf("======arr[%d]=\n [%s] \n", index, value)
				if strings.Index(value, "全国彩票") > 0 {
					reg := regexp.MustCompile(`全国共销售彩票([\d]+.[\d]+)\S+`)
					result := reg.FindStringSubmatch(value)
					if len(result) > 0 {
						caipiao.Total = result[1]
					}
				}
			}
		}
	}
	log.Println("GMOF_CaiPiao_HTMLSpider match caipiao.Total:%s", caipiao.Total)
	log.Printf("---------GMOF_CaiPiao_HTMLSpider Find() Finish----------")

	saveData_GMOF_CaiPiao_Month(caipiao)
	return err
}
