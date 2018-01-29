package main

import (
	"log"
	"regexp"
	"strings"
)

func person_home_pattern_full(justify string) (Birthday string, Home string) {
	Birthday = ""
	Home = ""
	if justify != "" {
		rows := strings.Split(justify, "。")
		//rows := strings.Split(justify, Convert2String("。", GB18030))
		for i := range rows {
			log.Printf("%d======home_pattern():%s", i, rows[i])
			Birthday, Home = person_home_pattern(rows[i])
			if Birthday != "" && Home != "" {
				return Birthday, Home
			}
		}
	}
	return Birthday, Home
}
func person_home_pattern(justify string) (Birthday string, Home string) {
	Birthday = ""
	Home = ""

	if Birthday == "" && Home == "" {
		exp := `([\d]+)年生于(\S+)，籍贯(\S+)市`
		reg := regexp.MustCompile(exp)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 0 {
			Birthday = result[1] + "年"
			Home = result[2] + "(" + result[3] + ")"
		}
	}
	if Birthday == "" && Home == "" {
		exp := `([\d]+)年([\d]+)月([\d]+)日出生于(\S+)，籍贯(\S+)`
		reg := regexp.MustCompile(exp)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 0 {
			Birthday = result[1] + "年" + result[2] + "月" + result[3] + "日"
			Home = result[4] + "(" + result[5] + ")"
		}
	}
	if Birthday == "" && Home == "" {
		exp := `([\d]+)年([\d]+)月生于(\S+)`
		reg := regexp.MustCompile(exp)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 0 {
			Birthday = result[1] + "年" + result[2] + "月"
			Home = result[3]
		}
	}
	if Birthday == "" && Home == "" {
		exp := `([\d]+)年([\d]+)月生，(\S+)人`
		reg := regexp.MustCompile(exp)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 0 {
			Birthday = result[1] + "年" + result[2] + "月"
			Home = result[3]
		}
	}
	if Birthday == "" && Home == "" {
		exp := `([\d]+)年([\d]+)月出生，(\S+)人`
		reg := regexp.MustCompile(exp)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 0 {
			Birthday = result[1] + "年" + result[2] + "月"
			Home = result[3]
		}
	}
	if Birthday == "" && Home == "" {
		exp := `([\d]+)年([\d]+)月出生，籍贯(\S+)`
		reg := regexp.MustCompile(exp)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 0 {
			Birthday = result[1] + "年" + result[2] + "月"
			Home = result[3]
		}
	}
	if Birthday == "" && Home == "" {
		exp := `([\d]+)年([\d]+)月([\d]+)日出生于(\S+)`
		reg := regexp.MustCompile(exp)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 0 {
			Birthday = result[1] + "年" + result[2] + "月" + result[3] + "日"
			Home = result[4]
		}
	}

	if Birthday == "" && Home == "" {
		exp := `([\d]+)年([\d]+)月([\d]+)日出生于(\S+)`
		reg := regexp.MustCompile(exp)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 0 {
			Birthday = result[1] + "年" + result[2] + "月" + result[3] + "日"
			Home = result[4]
		}
	}
	if Birthday == "" && Home == "" {
		exp := `([\d]+)年([\d]+)月([\d]+)日生于(\S+)`
		reg := regexp.MustCompile(exp)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 0 {
			Birthday = result[1] + "年" + result[2] + "月" + result[3] + "日"
			Home = result[4]
		}
	}

	if Birthday == "" && Home == "" {
		exp := `([\d]+)年([\d]+)月出生于(\S+)`
		reg := regexp.MustCompile(exp)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 3 {
			Birthday = result[1] + "年" + result[2] + "月"
			Home = result[3]
		}
	}
	if Birthday == "" && Home == "" {
		exp := `([\d]+)年([\d]+)月生于(\S+)`
		reg := regexp.MustCompile(exp)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 3 {
			Birthday = result[1] + "年" + result[2] + "月"
			Home = result[3]
		}
	}
	if Birthday == "" && Home == "" {
		exp := `([\d]+)年生于(\S+)`
		reg := regexp.MustCompile(exp)
		result := reg.FindStringSubmatch(justify)
		if len(result) > 0 {
			Birthday = result[1] + "年"
			Home = result[2]
		}
	}

	return Birthday, Home
}
