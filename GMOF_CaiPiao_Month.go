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

type GMOF_CaiPiao_Month struct {
	Title    string `json:"title"`
	AccId    string `json:"AccId"`
	Type     string `json:"Type"`
	Total    string `json:"Total"` //总计
	SPLY     string `json:"SPLY"`  //同比
	LP       string `json:"LP"`    //环比
	Url      string `json:"Url"`
	Attachid string `json:"Attachid"`
	Content  string `json:"content"`
}

type GMOF_CaiPiao_Months []GMOF_CaiPiao_Month

func Print_GMOF_CaiPiao(caipiao *GMOF_CaiPiao_Month) {
	Total := caipiao.Total
	LP := caipiao.LP
	SPLY := caipiao.SPLY
	attachid := caipiao.Attachid

	log.Printf("total:%s,LP:%s,SPLY:%s,attachid:%s", Total, LP, SPLY, attachid)
}

func NewGMOF_CaiPiao_Month() *GMOF_CaiPiao_Month {
	return &GMOF_CaiPiao_Month{
		Total:   "0",
		LP:      "0",
		SPLY:    "0",
		Content: "",
	}
}

func saveData_GMOF_CaiPiao_Month(caipiao *GMOF_CaiPiao_Month) {
	id := uuid.NewV4().String()
	itemcode := "GMOF_CaiPiao_Month"

	db, err := initPGConn()
	stmt, err := db.Prepare("insert into data_source_raw(id,itemcode,title,rawvalue,status)  VALUES($1,$2,$3,$4,$5) RETURNING id")
	checkDBErr(err)
	res, err := stmt.Exec(id, itemcode, caipiao.Title, caipiao.Total, "C")
	checkDBErr(err)
	affect, err := res.RowsAffected()
	fmt.Println("res affected:%s", affect)
	checkDBErr(err)
	db.Close()

}
