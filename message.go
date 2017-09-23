package main

type message struct {
	Type      string  `json:"type,omitempty"`
	Sender    string  `json:"sender,omitempty"`
	Message   string  `json:"message,omitempty"`
	Receiver  string  `json:"receiver,omitempty"`
	PositionX float64 `json:"positionX,omitempty"`
	PositionY float64 `json:"positionY,omitempty"`
}
