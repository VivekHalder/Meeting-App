package models

type Message struct {
	Type     string      `json:"message"`
	Payload  interface{} `json:"payload"`
	RoomID   string      `json:"roomID"`
	SenderID string      `json:"senderID"`
}
