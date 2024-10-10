package service

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"

	"n8body.dev/starb8ard-guilded/internal/guilded/dto"
	"n8body.dev/starb8ard-guilded/internal/message/service"
)

var GuildedWebSocketURL = "wss://www.guilded.gg/websocket/v1"

func CreateGuildedWebsocketConnection() *websocket.Conn {
	dialer := websocket.DefaultDialer

	log.Println("Establishing a connection to Guilded...")

	requestHeader := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", os.Getenv("ACCESS_TOKEN"))}}

	conn, _, err := dialer.Dial(GuildedWebSocketURL, requestHeader)
	if err != nil {
		log.Fatalln("Error connecting to Guilded:", err)
	}

	log.Println("Connected to Guilded WebSocket!")

	return conn
}

func CloseGuildedWebsocketConnection(conn *websocket.Conn) {
	log.Println("Closing the connection...")
	if err := conn.Close(); err != nil {
		log.Println("Error closing the connection:", err)
	}

	log.Println("Connection closed.")
}

func HandleWebsocketGuildedEvent(conn *websocket.Conn, message []byte) {
	event := dto.WebsocketGuildedEventDTO{}
	if err := json.Unmarshal(message, &event); err != nil {
		log.Println("Couldn't transform the following message to Guilded event, skipping:", err)
		return
	}

	switch event.T {
	case "ChatMessageCreated":
		service.HandleChatMessageCreated(conn, event.D)
	default:
		log.Println("Message not supported, skipping:", event.Op, event.T)
	}
}
