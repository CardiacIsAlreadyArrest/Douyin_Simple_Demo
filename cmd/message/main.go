package main

import (
	message "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/message/messageservice"
	"log"
)

func main() {
	svr := message.NewServer(new(MessageServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
