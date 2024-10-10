package dto

type ChatMessageCreatedEventDataDTO struct {
	ServerID string         `json:"serverId"`
	Message  ChatMessageDTO `json:"message"`
}
