package domain

type ChatWorkRequest struct {
	Webhook_event struct {
		Message_id  string
		Room_id     int32
		Account_id  int32
		Body        string
		Send_time   int32
		Update_time string
	}
}

type CallChatMessageId struct {
	Message_id string
}
