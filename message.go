package main

type message struct {
	Sender    string  `json:"sender"`
	Receiver  string  `json:"receiver"`
	Message   string  `json:"message"`
	PositionX float64 `json:"positionX"`
	PositionY float64 `json:"positionY"`
}
