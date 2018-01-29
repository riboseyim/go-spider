package test

import (
	"log"
	"regexp"
	"strings"
	"testing"
)

func Test_strings_Age(t *testing.T) {
	justify := "空气动力学家 1960年9月生于四川省绵阳市，籍贯四川中江。"
	Birthday, Home := person_home_pattern_full(justify)
	log.Printf("Birthday:%s,Home:%s", Birthday, Home)

	justify = "粒子物理学家  1946年生于湖北省武汉市，籍贯福建省福州市。1970年毕业于北京大学技术物理系，1984年获美国麻省理工学院博士学位。"
	Birthday, Home = person_home_pattern_full(justify)
	log.Printf("Birthday:%s,Home:%s", Birthday, Home)

	justify = " 无机化学家。中国科学技术大学教授。1967年7月23日出生于安徽省阜阳市，籍贯安徽安庆。1988年毕业于厦门大学化学系，1996年在中国科学技术大学应用化学系获博士学位。"
	Birthday, Home = person_home_pattern_full(justify)
	log.Printf("Birthday:%s,Home:%s", Birthday, Home)

	justify = "数学家 1963年9月生于江苏省靖江市。1982年毕业于中国科技大学数学系，1985年获中国科学院计算中心硕士学位，1989年获美国加州大学洛杉矶分校（UCLA）博士学位。"
	Birthday, Home = person_home_pattern_full(justify)
	log.Printf("Birthday:%s,Home:%s", Birthday, Home)

	justify = "女 天文学家 1951年12月生于重庆市万州，籍贯山东博兴。1975年毕业于南京理工大学光学仪器专业，1982年、1995年先后获中国科学院紫金山天文台硕士、博士学位。中国科学院国家天文台南京天文光学技术研究所研究员。"
	Birthday, Home = person_home_pattern_full(justify)
	log.Printf("Birthday:%s,Home:%s", Birthday, Home)
}

func person_home_pattern_full(justify string) (Birthday string, Home string) {
	Birthday = ""
	Home = ""
	rows := strings.Split(justify, "。")
	for i := range rows {
		Birthday, Home = person_home_pattern(rows[i])
		if Birthday != "" && Home != "" {
			return Birthday, Home
		}
	}
	return Birthday, Home
}
func person_home_pattern(justify string) (Birthday string, Home string) {
	Birthday = ""
	Home = ""

	log.Printf("======home_pattern():%s", justify)

	if Birthday == "" && Home == "" {
		reg := regexp.MustCompile(`([\d]+)年([\d]+)月([\d]+)日出生于(\S+)，籍贯(\S+)`)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 0 {
			Birthday = result[1] + "年" + result[2] + "月" + result[3] + "日"
			Home = result[4] + "(" + result[5] + ")"
		}
	}

	if Birthday == "" && Home == "" {
		reg := regexp.MustCompile(`([\d]+)年生于(\S+)，籍贯(\S+)市`)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 0 {
			Birthday = result[1] + "年"
			Home = result[2] + "(" + result[3] + ")"
		}
	}

	if Birthday == "" && Home == "" {
		reg := regexp.MustCompile(`([\d]+)年([\d]+)月生于(\S+)`)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 3 {
			Birthday = result[1] + "年" + result[2] + "月"
			Home = result[3]
		}
	}

	return Birthday, Home
}
