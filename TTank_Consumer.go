package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Shopify/sarama"
)

//consumer.go
func mainConsumer(partition int32, topic string) {
	kafka := newKafkaConsumer()
	defer kafka.Close()
	//注：开发环境中我们使用sarama.OffsetOldest，Kafka将从创建以来第一条消息开始发送。
	//在生产环境中切换为sarama.OffsetNewest,只会将最新生成的消息发送给我们。
	consumer, err := kafka.ConsumePartition(topic, partition, sarama.OffsetOldest)
	if err != nil {
		fmt.Printf("Kafka error: %s\n", err)
		os.Exit(-1)
	}
	go consumeEvents(consumer)

	fmt.Println("Press [enter] to exit consumer \n ")
	bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println("Terminating...")
}

func consumeEvents(consumer sarama.PartitionConsumer) {
	var msgVal []byte
	var log interface{}
	var logMap map[string]interface{}
	//	var caipiao *GMOF_CaiPiao_Month

	var err error
	for {
		//goruntine exec
		select {
		// blocking <- channel operator
		case err := <-consumer.Errors():
			fmt.Printf("Kafka error: %s\n", err)
		case msg := <-consumer.Messages():
			msgVal = msg.Value
			//
			if err = json.Unmarshal(msgVal, &log); err != nil {
				fmt.Printf("Failed parsing: %s", err)
			} else {
				logMap = log.(map[string]interface{})
				logType := logMap["Type"]
				fmt.Printf("Processing %s:\n%s\n", logMap["Type"], string(msgVal))

				switch logType {
				case "Add_GMOF_CaiPiao_Month":
					event := new(GMOF_CaiPiao_Month)
					if err = json.Unmarshal(msgVal, &event); err == nil {
						err = event.Process()
					}
				case "Add_GMOF_CaiPiao_List":
					event := new(GMOF_CaiPiao_List)
					if err = json.Unmarshal(msgVal, &event); err == nil {
						err = event.Process()
					}
				default:
					fmt.Println("Unknown command: ", logType)
				}
				if err != nil {
					fmt.Printf("Error processing: %s\n", err)
				} else {
					fmt.Println("---------------------")
					//fmt.Printf("%+v\n\n", *bankAccount)
				}
			}
		}
	}
}
