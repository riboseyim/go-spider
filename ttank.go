package main

import (
	"flag"
	"fmt"
)

const (
	VERSION = "-1.0-release-201801"
	AUTHOR  = "@RiboseYim"
)

func main() {

	act := flag.String("act", "producer", "Either: producer or consumer")
	task := flag.String("task", "", "GMOF_CaiPiao.eg")

	// partition := flag.String("partition", "0",
	// "Partition which the consumer program will be subscribing")

	flag.Parse()
	fmt.Printf("Welcome to [ TianGuan System %s ] Author:%s \n -act:%s \n\n", VERSION, AUTHOR, *act)

	switch *act {
	case "producer":
		// switch *task {
		// case "GMOF_CaiPiao_Default":
		// 	NewGMOFTask("GMOF_CaiPiao_Default")
		// default:
		// 	//log.Println("can not found task !")
		// }
		NewGMOFTask(*task)
	case "consumer":

	}

}
