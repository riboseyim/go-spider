package main

import (
	"flag"
	"log"
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

	log.Println("Welcome to [ TianGuan System %s ] Author:%s ] ", VERSION, AUTHOR)
	log.Println("Module:%s", MODULE)
	log.Println("-act:%s", *act)
	log.Println("-task:%s", *task)
	log.Println("-envflag:%s", *envflag)

	Exec_Module_Task_USATraffic(*act, *task, *envflag)

}
