package domain

type ChatWorkRequest struct {
	Webhook_event struct {
		Message_id  string
		Room_id     string
		Account_id  string
		Body        string
		Send_time   string
		Update_time string
	}
}

type CallChatMessageId struct {
	Message_id string
}
