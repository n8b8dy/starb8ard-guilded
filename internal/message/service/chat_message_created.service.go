package service

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"

	"n8body.dev/starb8ard-guilded/internal/message/dto"
)

func HandleChatMessageCreated(conn *websocket.Conn, data json.RawMessage) {
	createdMessage := dto.ChatMessageCreatedEventDataDTO{}
	if err := json.Unmarshal(data, &createdMessage); err != nil {
		log.Println("Error unmarshalling ChatMessageCreated data:", err)
		return
	}

	log.Println("Replying to a message created with content:", createdMessage.Message.Content)
}
