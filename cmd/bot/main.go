package main

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"

	"n8body.dev/starb8ard-guilded/internal/guilded/service"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file:", err)
	}

	conn := service.CreateGuildedWebsocketConnection()
	defer service.CloseGuildedWebsocketConnection(conn)

	for {
		log.Println("Reading a message...")

		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Couldn't read the message, skipping:", err)
			continue
		} else if messageType == websocket.CloseMessage {
			log.Println("Connection requested to be closed by Guilded.")
			break
		}

		go service.HandleWebsocketGuildedEvent(conn, message)
	}
}
