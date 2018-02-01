package main

import (
	"fmt"
)

type Echarts_HTML_Obj struct {
	Title        string
	SubTitle     string
	DataSource   string
	xArrays      []int
	Data         string
	TemplateName string
	TemplateFile string
	Target       string
	TopIndex     int
}

type USA_Traffic_Air struct {
	Country         string
	Year            int
	Passengers      int
	PassengersRatio float64
}

func CountRatio(e *USA_Traffic_Air, t []*USA_Traffic_Air) string {
	Total_Int := 0
	for i := range t {
		row := t[i]
		Total_Int += row.Passengers
	}
	Total := float64(Total_Int)

	if e != nil {
		Passengers := float64(e.Passengers)
		ratio := fmt.Sprintf("%.3f", Passengers/Total*100)
		//log.Printf("CountRatio :%v,%v,%v", ratio, Passengers, Total)
		return ratio
	}
	return "0"
}

type USA_Traffic_Air_ByPassengers []*USA_Traffic_Air

func (this USA_Traffic_Air_ByPassengers) Len() int {
	return len(this)
}

func (this USA_Traffic_Air_ByPassengers) Less(i, j int) bool {
	if this[i] != nil && this[j] != nil {
		return this[i].Passengers > this[j].Passengers
	}
	return false
}

func (this USA_Traffic_Air_ByPassengers) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func swap(a int, b int) (int, int) {
	return b, a
}

func partition(aris []int, begin int, end int) int {
	pvalue := aris[begin]
	i := begin
	j := begin + 1
	for j < end {
		if aris[j] < pvalue {
			i++
			aris[i], aris[j] = swap(aris[i], aris[j])
		}
		j++
	}
	aris[i], aris[begin] = swap(aris[i], aris[begin])
	return i
}

func quickSort(aris []int, begin int, end int) {
	if begin+1 < end {
		mid := partition(aris, begin, end)
		quickSort(aris, begin, mid)
		quickSort(aris, mid+1, end)
	}
}
