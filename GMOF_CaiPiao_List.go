package main

type GMOF_CaiPiao_List struct {
	AccId string `json:"AccId"`
	Title string `json:"title"`
	Type  string `json:"Type"`
	Url   string `json:"Url"`
}

type GMOF_CaiPiao_Lists []GMOF_CaiPiao_List

func NewGMOF_CaiPiao_List() *GMOF_CaiPiao_List {
	return &GMOF_CaiPiao_List{
		Url: "",
	}
}

func saveData_GMOF_CaiPiao_List(Id string, Title string, Type string, Url string, tt map[string]interface{}) {

	csvdata := [][]string{
		{Id, Title, Type, Url},
	}
	save_csv("./data/Data_GMOF_CaiPiao_List.csv", csvdata, true)

}
