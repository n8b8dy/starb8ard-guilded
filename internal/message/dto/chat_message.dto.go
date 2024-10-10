package dto

type ChatMessageDTO struct {
	Id   string `json:"id"`
	Type string `json:"type"`

	ServerID  string `json:"serverId,omitempty"`
	GroupID   string `json:"groupId,omitempty"`
	ChannelID string `json:"channelId"`

	Content string `json:"content"`

	CreatedAt string `json:"createdAt"`
	CreatedBy string `json:"createdBy"`
}
