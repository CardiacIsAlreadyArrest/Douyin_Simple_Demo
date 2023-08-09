package main

import (
	publish "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/publish/publishservice"
	"log"
)

func main() {
	svr := publish.NewServer(new(PublishServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
