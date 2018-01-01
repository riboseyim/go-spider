package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/celrenheit/spider"
	"github.com/celrenheit/spider/schedule"
)

var (
	// Ensure WikipediaHTMLSpider implements spider.Spider interface
	_ spider.Spider = (*WikipediaHTMLSpider)(nil)
	// Ensure WikipediaJSONSpider implements spider.Spider interface
	_ spider.Spider = (*WikipediaJSONSpider)(nil)
)
var now = time.Now()

func mainn_() {
	wikiHTMLSpider := &WikipediaHTMLSpider{"Redis_Labs"}

	spider.Add(schedule.Every(3*time.Second), wikiHTMLSpider)

	spider.Start()

	<-time.After(26 * time.Second)
}

type WikipediaHTMLSpider struct {
	Title string
}

func (w *WikipediaHTMLSpider) Setup(ctx *spider.Context) (*spider.Context, error) {
	url := fmt.Sprintf("https://en.wikipedia.org/wiki/%s", w.Title)
	return spider.NewHTTPContext("GET", url, nil)
}

func (w *WikipediaHTMLSpider) Spin(ctx *spider.Context) error {
	if _, err := ctx.DoRequest(); err != nil {
		return err
	}

	html, err := ctx.HTMLParser()
	if err != nil {
		return err
	}
	summary := html.Find("#mw-content-text p").First().Text()

	fmt.Println(summary)
	return nil
}

type WikipediaJSONSpider struct {
	Title string
}

func (w *WikipediaJSONSpider) Setup(ctx *spider.Context) (*spider.Context, error) {
	params := url.Values{}
	params.Add("titles", w.Title)
	url := fmt.Sprintf("http://en.wikipedia.org/w/api.php?format=json&action=query&prop=extracts&exintro=&explaintext=&%s", params.Encode())
	return spider.NewHTTPContext("GET", url, nil)
}

func (w *WikipediaJSONSpider) Spin(ctx *spider.Context) error {
	if _, err := ctx.DoRequest(); err != nil {
		return err
	}
	jsonparser, err := ctx.JSONParser()
	if err != nil {
		return err
	}
	pages, err := jsonparser.GetPath("query", "pages").Map()
	if err != nil {
		return err
	}
	for _, p := range pages {
		fmt.Println(p.(map[string]interface{})["extract"])
	}
	return nil
}
