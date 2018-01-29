package main

import (
	"fmt"
	"log"

	uuid "github.com/satori/go.uuid"
)

/*
1.2017-02 无数据
2.2012-04 及之前的样式不同

*/

type GMOF_Person_Region struct {
	Id        string `json:"Id"`
	Name      string `json:"Name"`
	Title     string `json:"Title"`
	Sex       string `json:"Sex"`
	Ethnic    string `json:"Ethnic"`
	Home      string `json:"Home"`
	Birthday  string `json:"Birthday"`
	Workday   string `json:"Workday"`
	Partyday  string `json:"Partyday"`
	Education string `json:"Education"`
	Summary   string `json:"Summary"`
	Resume    string `json:"Resume"`
	Sourceurl string `json:"Sourceurl"`
	LevelOne  string `json:"LevelOne"`
	LevelTwo  string `json:"LevelTwo"`
	Status    string `json:"Status"`
	Records   []string
}

func NewGMOF_Person_Region() *GMOF_Person_Region {
	return &GMOF_Person_Region{
		Title: "",
	}
}

func saveData_GMOF_Person_Region(person *GMOF_Person_Region) {
	log.Println("======saveData_GMOF_Person_Region()===========")
	log.Println("======Title:" + person.Title)

	id := uuid.NewV4().String()

	db, err := initPGConn()
	stmt, err := db.Prepare("insert into person(id,name,position,sex,ethnic,home,Birthday,education,workday,partyday,summary,resume,sourceurl,status)  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) RETURNING id")
	checkDBErr(err)
	if len(person.Ethnic) > 50 {
		person.Ethnic = GoSubstr(person.Ethnic, 0, 50)
	}
	if len(person.Home) > 50 {
		person.Home = GoSubstr(person.Home, 0, 50)
	}
	if len(person.Workday) > 100 {
		person.Workday = GoSubstr(person.Workday, 0, 100)
	}
	if len(person.Partyday) > 100 {
		person.Partyday = GoSubstr(person.Partyday, 0, 100)
	}
	if len(person.Birthday) > 100 {
		person.Birthday = GoSubstr(person.Birthday, 0, 100)
	}
	if len(person.Education) > 100 {
		person.Education = GoSubstr(person.Education, 0, 200)
	}

	res, err := stmt.Exec(id, person.Name, person.Title, person.Sex, person.Ethnic, person.Home, person.Birthday, person.Education, person.Workday, person.Partyday, person.Summary, person.Resume, person.Sourceurl, "NEW")
	checkDBErr(err)
	affect, err := res.RowsAffected()
	fmt.Println("res affected:%s", affect)

	//异常处理仍需优化，20180108
	if err != nil {
		panic(err)
		log.Printf("=====Error:%s", person.Sourceurl)
	}

	db.Close()
}
