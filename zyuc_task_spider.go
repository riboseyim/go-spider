package main

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/celrenheit/spider"
)

type ZYUC_Task_Info_HTMLSpider struct {
	title string `json:"title"`
	url   string `json:"url"`
	desc  string `json:"desc"`
}

func NewZYUC_Task_Info_HTMLSpider() *ZYUC_Task_Info_HTMLSpider {
	return &ZYUC_Task_Info_HTMLSpider{
		desc: "",
	}
}

func (w *ZYUC_Task_Info_HTMLSpider) Setup(ctx *spider.Context) (*spider.Context, error) {
	return spider.NewHTTPContext("POST", w.url, nil)
}

func (w *ZYUC_Task_Info_HTMLSpider) Spin(ctx *spider.Context) error {
	if _, err := ctx.DoRequest(); err != nil {
		return err
	}

	html, err := ctx.HTMLParser()
	if err != nil {
		return err
	}

	//<title></title>
	title := html.Find("title").Eq(0).Text()
	ProjectName, _ := html.Find("#projectname").Eq(0).Attr("value")
	TaskID, _ := html.Find("#taskID").Eq(0).Attr("value")
	myHiddenStatus4BaseInfo, _ := html.Find("#myHiddenStatus4BaseInfo").Eq(0).Attr("value") //20

	//timePromised, _ := html.Find("#base_timePromised").Eq(0).Attr("value") //承诺完成时间
	//*[@id="base_timePromised"]
	//#base_timePromised

	timePromised := ""
	t, _ := html.Find("#base_timePromised").Html()
	log.Println("base_timePromised=====%s", t)

	html.Find("#base_timePromised").Each(func(i int, s *goquery.Selection) {

		t, _ := s.Html()
		log.Println("base_timePromised=====%s", t)

		s.Find("tr").Eq(6).Each(func(i int, s *goquery.Selection) {
			log.Println("tr=====%v", s)
			timePromised = s.Find("td").Eq(3).Text()
		})
	})

	log.Println("===1 base_timePromised:%s", timePromised)
	log.Println("=== %s \n %s,%s,%s,%s \n", title, ProjectName, TaskID, timePromised, myHiddenStatus4BaseInfo)

	task := NewZYUC_Task_Info()
	task.TaskID = TaskID
	task.ProjectName = ProjectName
	task.PromiseTime = timePromised
	task.Status = myHiddenStatus4BaseInfo
	log.Println("===2 base_timePromised:%s", task.PromiseTime)
	SaveCSV_ZYUC_Task_Info(task)
	return err
}
