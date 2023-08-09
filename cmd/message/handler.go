package main

import (
	"context"
	message "github.com/Yra-A/Douyin_Simple_Demo/kitex_gen/message"
)

// MessageServiceImpl implements the last service interface defined in the IDL.
type MessageServiceImpl struct{}

// MessageAction implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageAction(ctx context.Context, req *message.MessageActionRequest) (resp *message.MessageActionResponse, err error) {
	// TODO: Your code here...
	return
}

// MessageChat implements the MessageServiceImpl interface.
func (s *MessageServiceImpl) MessageChat(ctx context.Context, req *message.MessageChatRequest) (resp *message.MessageChatResponse, err error) {
	// TODO: Your code here...
	return
}
