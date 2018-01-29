package main

import (
	"flag"
	"fmt"
)

const (
	VERSION = "-1.0-release-201801"
	MODULE  = "Economist"
	AUTHOR  = "@RiboseYim"
)

func main() {
	act := flag.String("act", "", "[producer]")
	task := flag.String("task", "", "[build-echarts]")
	envflag := flag.String("envflag", "", "[0:test;1:production]")

	flag.Parse()

	if *envflag == "" {
		*envflag = "1"
	}

	fmt.Printf("Welcome to [ TianGuan System %s ] Author:%s ] \n", VERSION, AUTHOR)
	fmt.Printf("Module:%s \n\n", MODULE)
	fmt.Printf("-act:%s \n", *act)
	fmt.Printf("-task:%s \n", *task)
	fmt.Printf("-envflag:%s \n", *envflag)

	switch *act {
	case "producer":
		Exec_Module_Task_USATraffic(*task, *envflag)
	case "consumer":
		//

	}

}
