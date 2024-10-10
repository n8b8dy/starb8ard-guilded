package dto

import "encoding/json"

type WebsocketGuildedEventDTO struct {
	Op int             `json:"op"`
	T  string          `json:"t,omitempty"`
	S  string          `json:"s,omitempty"`
	D  json.RawMessage `json:"d,omitempty"`
}
