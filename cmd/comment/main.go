package main

import (
	comment "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/comment/commentservice"
	"log"
)

func main() {
	svr := comment.NewServer(new(CommentServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
