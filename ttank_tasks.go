package main

import (
	"log"
)

func NewGMOFTask(task string) {
	if task == "GMOF_Person_Test" {
		Handle_GMOF_Person_Region_Test()
	} else if task == "GMOF_Person_List" {
		Handle_GMOF_Person_Region_BatchTask()
	} else if task == "GMOF_CaiPiao_Test" {
		Handle_GMOF_CaiPiao_Month_Test()
	} else if task == "GMOF_CaiPiao_Month" {
		Handle_GMOF_CaiPiao_Month_BatchTask()
	} else if task == "GMOF_CaiPiao_List" {
		Handle_GMOF_CaiPiao_List_BatchTask()
	} else if task == "GMOF_CDC_Epide_Month_Test" {
		Handle_GMOF_CDC_Epide_Month_Test()
	} else if task == "GMOF_Casad_List" {
		Handle_GMOF_Casad_List_Task()
	} else if task == "GMOF_Casad_ReQuery" {
		Handle_GMOF_Casad_ReQuery_Task()
	} else if task == "GMOF_Casad_Test" {
		Handle_GMOF_Casad_Info_Test()
	} else if task == "ZYUC_Task_Info_Collect" {
		Handle_ZYUC_Task_Info_Collect()
	} else {
		log.Println("Can not match this taskId:[%s]", task)
	}

}
