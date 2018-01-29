package main

type GMOF_CDC_Epide_List struct {
	AccId string `json:"AccId"`
	Title string `json:"title"`
	Type  string `json:"Type"`
	Url   string `json:"Url"`
}

type GMOF_CDC_Epide_Lists []GMOF_CDC_Epide_List

func NewGMOF_CDC_Epide_List() *GMOF_CDC_Epide_List {
	return &GMOF_CDC_Epide_List{
		Url: "",
	}
}

func saveData_GMOF_CDC_Epide_List(event *GMOF_CDC_Epide_List) {

	csvdata := [][]string{
		{event.AccId, event.Title, event.Type, event.Url},
	}

	save_csv("./data/Data_GMOF_CDC_Epide_List.csv", csvdata, true)

}
