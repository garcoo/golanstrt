package chatwork

import (
	"net/http"

	"../../application"
	"../../domain"
	"github.com/ant0ine/go-json-rest/rest"
)

func ContactChatWork(w rest.ResponseWriter, req *rest.Request) {
	input := domain.ChatWorkRequest{}
	err := req.DecodeJsonPayload(&input)

	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(&domain.CallChatMessageId{
		application.ContactBot(input),
	})
}
