package main

import (
	"bytes"
	"log"
	"os"
	"text/template"
)

type Person struct {
	UserName string
}

func main() {

	t := template.New("usa-traffic.tmpl")        //创建一个模板
	t, _ = t.ParseFiles("tmpl/usa-traffic.tmpl") //解析模板文件
	p := Person{UserName: "Astaxie"}             //获取当前用户信息

	var tpl bytes.Buffer
	t.Execute(&tpl, p)
	result := tpl.String()
	log.Printf(result)

	f, _ := os.Create("tmpl/usa-traffic.js")
	f.WriteString(result)
}
