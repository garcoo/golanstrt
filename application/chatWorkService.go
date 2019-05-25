package application

import (
	"fmt"

	"../domain"
	"../infrastructure"
)

func ContactBot(request domain.ChatWorkRequest) domain.CallChatWorkResponse {

	var users = infrastructure.FindOneByChatAccountId(request.Webhook_event.Account_id)
	fmt.Println(users[0].User_nickname)
	return domain.CallChatWorkResponse{Message_id: users[0].User_nickname}
}
