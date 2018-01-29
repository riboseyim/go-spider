package main

import (
	"log"
)

/*
1.2017-02 无数据
2.2012-04 及之前的样式不同

*/

type GMOF_CDC_Epide_Month struct {
	Title     string `json:"Title"`
	Name      string `json:"Name"`
	Incidence string `json:"Incidence"`
	Death     string `json:"Death"`
	Content   string `json:"content"`
}

type GMOF_CDC_Epide_Months []GMOF_CDC_Epide_Month

func Print_GMOF_CDC_Epide(CDC_Epide *GMOF_CDC_Epide_Month) {

	log.Printf("Print_GMOF_CDC_Epide()------")
}

func NewGMOF_CDC_Epide_Month() *GMOF_CDC_Epide_Month {
	return &GMOF_CDC_Epide_Month{
		Title:     "",
		Name:      "",
		Incidence: "0",
		Death:     "0",
		Content:   "",
	}
}

func saveData_GMOF_CDC_Epide_Month(CDC_Epide *GMOF_CDC_Epide_Month) {
	//	id := uuid.NewV4().String()
	//	itemcode := "GMOF_CDC_Epide_Month"

	// db, err := initPGConn()
	// stmt, err := db.Prepare("insert into data_source_raw(id,itemcode,title,rawvalue,status)  VALUES($1,$2,$3,$4,$5) RETURNING id")
	// checkDBErr(err)
	//	res, err := stmt.Exec(id, itemcode, CDC_Epide.Title, CDC_Epide.Total, "C")
	// checkDBErr(err)
	// affect, err := res.RowsAffected()
	// fmt.Println("res affected:%s", affect)
	// checkDBErr(err)
	// db.Close()

}
