package main

func Exec_Module_Task_USATraffic(act string, task string, envflag string) {
	if act == "build-echarts" {
		Exec_Module_Task_Echarts(task, envflag)
	} else {
		Exec_Module_Task_Echarts(task, envflag)
	}
}
