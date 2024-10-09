package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
)

var GuildedWebSocketURL = "wss://www.guilded.gg/websocket/v1"

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	dialer := websocket.DefaultDialer

	log.Println("Establishing a connection to Guilded...")

	requestHeader := http.Header{"Authorization": {fmt.Sprintf("Bearer %v", os.Getenv("ACCESS_TOKEN"))}}
	conn, _, err := dialer.Dial(GuildedWebSocketURL, requestHeader)
	if err != nil {
		log.Fatalln("Error connecting to Guilded:", err)
	}

	defer func() {
		log.Println("Closing the connection...")

		if err := conn.Close(); err != nil {
			log.Println("Error closing the connection:", err)
		}

		log.Println("Connection closed.")
	}()

	log.Println("Connected to Guilded WebSocket!")

	for {
		log.Println("Reading a message...")
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Couldn't read the message, skipping:", err)
			continue
		} else if messageType == websocket.CloseMessage {
			break
		}

		go handleGuildedEvent(conn, message)
	}
}

type ChatMessageDTO struct {
	Id      string `json:"id"`
	Content string `json:"content"`
}

type ChatMessageCreatedDataDTO struct {
	ServerId string         `json:"serverId"`
	Message  ChatMessageDTO `json:"message"`
}

type GuildedEventDTO struct {
	Op int             `json:"op"`
	T  string          `json:"t"`
	S  string          `json:"s"`
	D  json.RawMessage `json:"d"`
}

func handleChatMessageCreated(conn *websocket.Conn, data json.RawMessage) {
	createdMessage := ChatMessageCreatedDataDTO{}
	if err := json.Unmarshal(data, &createdMessage); err != nil {
		log.Println("Error unmarshalling ChatMessageCreated data:", err)
		return
	}

	log.Println("Replying to a message created with content:", createdMessage.Message.Content)
}

func handleGuildedEvent(conn *websocket.Conn, message []byte) {
	event := GuildedEventDTO{}
	if err := json.Unmarshal(message, &event); err != nil {
		log.Println("Couldn't transform the following message to guilded event, skipping:", err)
		return
	}

	switch event.T {
	case "ChatMessageCreated":
		handleChatMessageCreated(conn, event.D)
	default:
		log.Println("Message not supported, skipping:", event.Op, event.T)
	}
}
