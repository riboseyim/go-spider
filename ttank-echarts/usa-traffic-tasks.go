package main

import (
	"bytes"
	"container/list"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"text/template"
)

func Exec_Module_Task_Echarts(task string, envflag string) {
	log.Println("----- Exec_Module_Task_Echarts() begin ")
	if task == "usa-traffic" {
		obj := &Echarts_HTML_Obj{
			Title:        "美国航空入境旅客(出发地)变化情况 @RiboseYim",
			SubTitle:     "数据来源：美国交通部 · 单位：千人次 \\n https://riboseyim.github.io",
			Data:         "",
			TemplateName: "usa-traffic.tmpl",
			TemplateFile: "template/usa-traffic.tmpl",
			Target:       "./build/usa-traffic.js",
		}
		if envflag == "0" {
			obj.xArrays = []int{2001, 2002}
			obj.DataSource = "./data/usa-traffic-data-test.dat"
			obj.TopIndex = 2
			Build_USATraffic_HTML_ECharts(obj)
		} else {
			obj.xArrays = []int{1975, 1980, 1985, 1990, 1991, 1992, 1993, 1994, 1995, 1996, 1997, 1998, 1999, 2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014, 2015}
			obj.DataSource = "./data/usa-traffic-data-all.dat"
			obj.TopIndex = 15
			Build_USATraffic_HTML_ECharts(obj)
		}

	} else if task == "eu-ip-2015" {
		obj := &Echarts_HTML_Obj{
			Title:        "世界主要国家专利数量变化情况 @RiboseYim",
			SubTitle:     "数据来源：欧洲专利局 · 单位：文献索引数 \\n https://riboseyim.github.io",
			Data:         "",
			TemplateName: "eu-ip.tmpl",
			TemplateFile: "template/eu-ip.tmpl",
			Target:       "./build/eu-ip-2015.js",
		}

		obj.xArrays = []int{2015}
		obj.TopIndex = 20
		obj.DataSource = "./data/eu-ip-data-test.dat"
		Build_USATraffic_HTML_ECharts(obj)

	}

	log.Println("----- Exec_Module_Task_Echarts() end ")
}

func Build_USATraffic_HTML_ECharts(obj *Echarts_HTML_Obj) {
	log.Println("----- Build_USATraffic_HTML_ECharts() ------")
	dataMap, dataList := Load_USATraffic_Data(obj.DataSource, obj.xArrays)
	log.Println("----- Load_USATraffic_Data() finish------")
	obj.Data = Build_USATraffic_ECharts_JS(dataMap, dataList, obj.xArrays, obj.TopIndex)
	log.Println("----- Build_USATraffic_ECharts_JS() finish------")
	Drawing_HTML_Template(obj)
	log.Println("----- Drawing_HTML_Template() finish ------")
}

func Drawing_HTML_Template(obj *Echarts_HTML_Obj) {
	log.Println("----- Drawing_HTML_Template() begin")
	log.Println("----- Drawing_HTML_Template() Title:%s", obj.Title)
	log.Println("----- Drawing_HTML_Template() Template:%s,%s", obj.TemplateName, obj.TemplateFile)
	log.Println("----- Drawing_HTML_Template() Target:%s", obj.Target)
	//log.Println("-----Drawing_HTML_Template() Data:%s", obj.Data)

	t := template.New(obj.TemplateName)   //创建一个模板
	t, _ = t.ParseFiles(obj.TemplateFile) //解析模板文件

	var tpl bytes.Buffer
	t.Execute(&tpl, obj)
	result := tpl.String()
	//log.Println("-----Template Output:\n %s", result)

	f, _ := os.Create(obj.Target)
	f.WriteString(result)
}

func Build_USATraffic_ECharts_JS(dataMap map[string]map[int]string, dataList *list.List, USA_Traffic_Years []int, TopIndex int) string {
	html_js := ""

	yearMap := make(map[int][]*USA_Traffic_Air, len(USA_Traffic_Years))

	for j := range USA_Traffic_Years {
		year_this := USA_Traffic_Years[j]
		country_array := make([]*USA_Traffic_Air, len(dataMap))
		year_array_rownum := 0
		for e := dataList.Front(); e != nil; e = e.Next() {
			cell, ok := (e.Value).(*USA_Traffic_Air)
			if ok && cell != nil {
				if cell.Year == year_this {
					//log.Println("-----add : *USA_Traffic_Air %s", cell)
					country_array[year_array_rownum] = cell
					year_array_rownum++
				}
			} else {
				log.Println("-----can not match this stuct : *USA_Traffic_Air")
			}
		}
		sort.Sort(USA_Traffic_Air_ByPassengers(country_array))
		yearMap[year_this] = country_array
	}

	//--------------------------------------------------------

	for i := range USA_Traffic_Years {
		year := USA_Traffic_Years[i]
		country_array := yearMap[year]

		year_js := ToString("{ \n  \"time\":", year, ",\"data\":[ \n")

		for k := range country_array {
			cell := country_array[k]
			country_js := ""
			if k < TopIndex {
				if cell != nil {
					Country := cell.Country
					CountryCName := ToCountryCN(Country)
					Passengers := cell.Passengers
					Ratio := CountRatio(cell, country_array)
					this_row := ""
					//log.Println("-----add year_array cell:%s,%s,%s,%s", Country, cell.Year, Passengers, CountryCName)
					if k < TopIndex-1 {
						this_row = ToString(" {\"name\":\"", Country, "\",\"value\":[", Passengers, ",", Ratio, ",\""+CountryCName+"\"]},")
					} else {
						this_row = ToString(" {\"name\":\"", Country, "\",\"value\":[", Passengers, ",", Ratio, ",\""+CountryCName+"\"]}")
					}
					//log.Println("-----add year_array js:%s", this_row)
					country_js += this_row
				} else {
					log.Println("cell is nil :\n %s ", cell)
				}
				year_js += country_js + "\n"
			}
		}

		if i < len(yearMap)-1 {
			html_js += year_js + "  ]\n},\n"
		} else {
			html_js += year_js + "  ]\n}\n"
		}

		i++
	}

	//log.Println("html_js:\n %s ", html_js)
	return html_js
}

func Load_USATraffic_Data(filename string, USA_Traffic_Years []int) (map[string]map[int]string, *list.List) {
	target_data := make(map[string]map[int]string, 50)
	source_data := read_csv_normal_2(filename)

	list := list.New()

	for i := range source_data {
		if source_data[i] != "" {
			rows := strings.Split(source_data[i], ",")
			//log.Println("-----query rows:%s", rows)
			country := ""
			countryMap := make(map[int]string)

			if len(rows) > 0 {
				country = rows[1]
				//log.Println("-----pre country:%s,len(rows)", country, len(rows))
				for j := 2; j < len(rows); j++ {
					Passengers_str := rows[j]
					if country != "" && Passengers_str != "" {
						//log.Println("-----USA_Traffic_Years:%s,j:%s", USA_Traffic_Years, j)
						year := USA_Traffic_Years[j-2]
						if rows[j] == "N" {
							Passengers_str = "0"
						}
						countryMap[year] = Passengers_str
						Passengers_64, _ := strconv.ParseInt(Passengers_str, 10, 64)
						Passengers := int(Passengers_64)
						cell := &USA_Traffic_Air{Country: country, Year: year, Passengers: Passengers}
						list.PushBack(cell)
						//	log.Println("-----list push country:%s,year:%s,passengers:%s", country, year, Passengers)
					}
				}
			}
			if country != "" {
				target_data[country] = countryMap
			}
		}
	}
	return target_data, list
}

func ToCountryCN(eng string) string {
	//cc := map[string]string{"a": "", "b": ""}
	cc := map[string]string{
		"Canada":               "加拿大",
		"Mexico":               "墨西哥",
		"United Kingdom":       "英国",
		"Japan":                "日本",
		"Germany":              "德国",
		"China":                "中国（含台湾）",
		"Hong Kong":            "中国（香港）",
		"France":               "法国",
		"Dominican Republic":   "多米尼加",
		"South Korea":          "韩国",
		"Brazil":               "巴西",
		"Netherlands":          "荷兰",
		"Jamaica":              "牙买加",
		"Italy":                "意大利",
		"Colombia":             "哥伦比亚",
		"Spain":                "西班牙",
		"Australia":            "澳大利亚",
		"Panama":               "巴拿马",
		"The Bahamas":          "巴哈马群岛",
		"Ireland":              "爱尔兰",
		"Switzerland":          "瑞士",
		"Israel":               "以色列",
		"Belgium":              "比利时",
		"Haiti":                "海地",
		"Denmark":              "丹麦",
		"Venezuela":            "委内瑞拉",
		"Cayman Islands":       "开曼群岛",
		"Philippines":          "菲律宾",
		"Bermuda":              "百慕大群岛",
		"Barbados":             "巴巴多斯岛",
		"Greece":               "希腊",
		"Netherlands Antilles": "荷属安的列斯",
	}

	return cc[eng]
}

func ToString(args ...interface{}) string {
	result := ""
	for _, arg := range args {
		switch val := arg.(type) {
		case int:
			result += strconv.Itoa(val)
		case string:
			result += val
		case float64:
			result += strconv.FormatFloat(val, 'g', 1, 64)
		}
	}
	return result
}
