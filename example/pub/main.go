package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nsqio/go-nsq"
)

var producer *nsq.Producer

func initProducer(addr string) (err error) {
	config := nsq.NewConfig()
	producer, err = nsq.NewProducer(addr, config)
	if err != nil {
		fmt.Printf("created product failed, err :%v \n", err.Error())
		return err
	}
	return nil
}

func main() {
	nsqAddress := "127.0.0.1:4150"
	err := initProducer(nsqAddress)
	if err != nil {
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("readd err 1")
			continue
		}
		data = strings.TrimSpace(data)
		fmt.Println("data: ==== ", data)
		if strings.ToUpper(data) == "Q" {
			break
		}
		err = producer.Publish("topic_demo", []byte(data))
		if err != nil {
			fmt.Printf("publish msg to nsq faild %v\n", err.Error())
			continue
		}
	}
}
