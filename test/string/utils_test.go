package test

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"testing"
)

func Test_strings_Sub1(t *testing.T) {
	title := "彭清华 简历 - 人民网 地方领导资料库"
	names := strings.Split(title, " ")
	if len(names) > 0 {
		name := names[0]
		log.Println("====Name:%s", name)
	}

}

func Test_strings_Sub2(t *testing.T) {
	title := "彭清华 简历 - 人民网 地方领导资料库"
	name := GoSubstr(title, 0, strings.Index(title, "简历"))
	log.Println("====Name:%s", name)
}

func GoSubstr(str string, start int, end int) string {
	rs := []rune(str)
	length := len(rs)

	if start < 0 || start > length {
		return ""
	}

	if end < 0 || end > length {
		return ""
	}
	return string(rs[start:end])
}

func Test_strings_Array(t *testing.T) {
	i := 0
	rows := make([]string, 10)

	rows[i] = "1"
	rows[i] = "1"
	rows[i] = "1"

	for index, value := range rows {
		fmt.Printf("arr[%d]=%s \n", index, value)
	}
}

func Test_strings_Split(t *testing.T) {
	Content := "a,b，c"
	rows := strings.Split(Content, ",")

	for index, value := range rows {
		fmt.Printf("arr[%d]=%s \n", index, value)
	}
}

func Test_strings_rows(t *testing.T) {
	Content := " 2017年12月22日　来源：综合司　　一、全国彩票销售情况　　11月份，全国共销售彩票385.55亿元，比上年同期（简称“同比”）增加40.72亿元，增长11.8%。其中，福利彩票机构销售194.02亿元，同比增加15.41亿元，增长8.6%；体育彩票机构销售191.52亿元，同比增加25.31亿元，增长15.2%。　　1-11月累计，全国共销售彩票3869.87亿元，同比增加289.40亿元，增长8.1%。其中，福利彩票机构销售1963.27亿元，同比增加98.20亿元，增长5.3%；体育彩票机构销售1906.60亿元，同比增加191.20亿元，增长11.1%。　　　　二、分类型彩票销售情况　　11月份，乐透数字型彩票销售235.85亿元，同比增加18.89亿元，增长8.7%；竞猜型彩票销售90.54亿元，同比增加21.52亿元，增长31.2%；即开型彩票销售19.47亿元，同比减少1.81亿元，下降8.5%；视频型彩票销售39.57亿元，同比增加2.19亿元，增长5.9%；基诺型彩票[1]销售0.12亿元，同比减少0.06亿元，下降33.7%。11月份，乐透数字型、竞猜型、即开型、视频型、基诺型彩票销售量分别占彩票销售总量的61.1%、23.5%、5.0%、10.3%、0.1%。　　1-11月累计，乐透数字型彩票销售2377.90亿元，同比增加163.16亿元，增长7.4%；竞猜型彩票销售842.66亿元，同比增加141.78亿元，增长20.2%；即开型彩票销售225.83亿元，同比减少32.22亿元，下降12.5%；视频型彩票销售421.79亿元，同比增加17.48亿元，增长4.3%；基诺型彩票销售1.70亿元，同比减少0.80亿元，下降31.9%。1-11月乐透数字型、竞猜型、即开型、视频型和基诺型彩票销售量分别占彩票销售总量的61.4%、21.8%、5.8%、10.9%、0.1%。　　　　三、分地区彩票销售情况　　11月份，与上年同期相比，全国共有27个省份彩票销售量出现增长。其中，江苏、安徽、广东、广西和湖南增加额比较多，同比分别增加6.49亿元、4.15亿元、4.06亿元、3.78亿元和3.69亿元。　　1-11月累计，与上年同期相比，全国共有27个省份彩票销售量出现增长。其中，湖北、福建、江西、重庆和广东增加额较多，同比分别增加26.47亿元、26.35亿元、25.25亿元、23.95亿元和23.45亿元。　　各级彩票机构要密切跟踪分析新情况新问题，切实加强彩票发行销售工作，确保市场平稳运行。各级财政部门要进一步加强监管，积极创造良好的外部环境，维护市场正常秩序，促进彩票事业健康发展。　　　　[1] 从2015年1月起，基诺型彩票销售量单独统计，不再计入乐透数字型彩票销售量当中。"
	log.Println("Test_strings")
	fmt.Printf("----- %q ---------\n", strings.Split(Content, "，"))

	rows := strings.Split(Content, "。")

	for index, value := range rows {
		fmt.Printf("arr[%d]=%s \n", index, value)

		if strings.Index(value, "全国彩票") > 0 {
			//log.Printf("=======Matched Row =======%s", value)
			reg := regexp.MustCompile(`全国共销售彩票([\d]+.[\d]+)\S+`)
			result := reg.FindStringSubmatch(value)
			log.Printf("----Parser_GMOF_CaiPiao()----\n %q \n ------ \n", result)
		}
	}

}
