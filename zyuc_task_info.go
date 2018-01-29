package main

import (
	"fmt"
	"log"
)

type ZYUC_Task_Info struct {
	TaskID      string `json:"TaskID"`
	ProjectName string `json:"ProjectName"`
	PromiseTime string `json:"PromiseTime"`
	Status      string `json:"Status"`
}

type ZYUC_Task_Infos []ZYUC_Task_Info

func NewZYUC_Task_Info() *ZYUC_Task_Info {
	return &ZYUC_Task_Info{
		TaskID: "",
	}
}

func SaveCSV_ZYUC_Task_Info(task *ZYUC_Task_Info) {
	log.Println("===3 base_timePromised:%s ", task.PromiseTime)
	csvdata := [][]string{
		{task.ProjectName, task.TaskID, task.PromiseTime, task.Status},
	}
	save_csv("./zyuc/Task_List_New.cfg", csvdata, true)
}

func query_ZYUC_Task_Info_NoHome() (result map[string]string) {
	db, err := initPGConn()
	//查询数据
	sql := "SELECT id,name,home,sourceurl FROM person where batchnum is null and home='' and sourceurl !='' "
	log.Println(sql)
	rows, err := db.Query(sql)
	checkDBErr(err)

	persons := make(map[string]string)
	for rows.Next() {
		var id string
		var name string
		var home string
		var sourceurl string

		err = rows.Scan(&id, &name, &home, &sourceurl)
		checkDBErr(err)

		persons[id] = sourceurl

	}

	db.Close()
	return persons
}

func Delete_ZYUC_Task_Info(sourceurl string) {

	db, err := initPGConn()
	stmt, err := db.Prepare("delete from person where batchnum is null and sourceurl=$1  ")
	checkDBErr(err)

	res, err := stmt.Exec(sourceurl)
	checkDBErr(err)
	affect, err := res.RowsAffected()
	fmt.Println("res affected:%s", affect)

	//异常处理仍需优化，20180108
	if err != nil {
		panic(err)
		log.Printf("=====Error:%s", sourceurl)
	}

	db.Close()
}

func Save_ZYUC_Task_Info(person ZYUC_Task_Info) {
	// log.Println("======saveData_ZYUC_Task_Info()===========")
	//
	// id := uuid.NewV4().String()

	// db, err := initPGConn()
	// stmt, err := db.Prepare("insert into person(id,name,position,sex,ethnic,home,Birthday,education,workday,partyday,summary,resume,sourceurl,status)  VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14) RETURNING id")
	// checkDBErr(err)
	//
	// res, err := stmt.Exec(id, person.Name, person.Title, person.Sex, person.Ethnic, person.Home, person.Birthday, person.Education, person.Workday, person.Partyday, person.Summary, person.Resume, person.Sourceurl, "NEW")
	// checkDBErr(err)
	// affect, err := res.RowsAffected()
	// fmt.Println("res affected:%s", affect)
	//
	// //异常处理仍需优化，20180108
	// if err != nil {
	// 	panic(err)
	// 	log.Printf("=====Error:%s", person.Sourceurl)
	// }
	//
	// db.Close()
}
