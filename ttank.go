package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
)

func main() {

	act := flag.String("act", "producer", "Either: producer or consumer")
	partition := flag.String("partition", "0",
		"Partition which the consumer program will be subscribing")
	taskname := flag.String("taskname", "", "GMOF_CaiPiao.eg")
	topicname := flag.String("topicname", topic_ttank_gmof_caipiao_list, "")

	flag.Parse()
	fmt.Printf("Welcome to ttank-microservice : %s \n\n", *act)

	switch *act {
	case "producer":
		switch *taskname {
		case "GMOF_CaiPiao_Default":
			NewTTankTask("GMOF_CaiPiao_Default")
		case "GMOF_CaiPiao_CSV":
			NewTTankTask("GMOF_CaiPiao_CSV")
		case "GMOF_CaiPiao_List":
			NewTTankTask("GMOF_CaiPiao_List")
		default:
			log.Println("can not found task !")
		}
	case "consumer":
		if part32int, err := strconv.ParseInt(*partition, 10, 32); err == nil {
			switch *topicname {
			case topic_ttank_gmof_caipiao_month:
				mainConsumer(int32(part32int), topic_ttank_gmof_caipiao_month)
			case topic_ttank_gmof_caipiao_list:
				mainConsumer(int32(part32int), topic_ttank_gmof_caipiao_list)
			default:
				log.Println("can not found topicname !")
			}

		}
	}

}
