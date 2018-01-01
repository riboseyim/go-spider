package main

import "log"

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

func saveData_GMOF_CaiPiao_Month(Id string, Title string, Total string, SPLY string, LP string, Url string, Attachid string, tt map[string]interface{}) {
	log.Println("======saveData_GMOF_CaiPiao_Month()===========")
	log.Println("======total:" + Total)

	csvdata := [][]string{
		{Id, Title, Total, SPLY, LP, Url, Attachid},
	}
	save_csv("./data/Data_GMOF_CaiPiao_Month.csv", csvdata, true)
}

/*
func updateAccount(id string, data map[string]interface{}) (*BankAccount, error) {
	cmd := Redis.HMSet(id, data)

	if err := cmd.Err(); err != nil {
		return nil, err
	} else {
		return FetchAccount(id)
	}
}
*/
