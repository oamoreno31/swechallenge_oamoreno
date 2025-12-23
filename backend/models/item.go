package models

type Item struct {
	ID          int64  `json:"ID"`
	Ticker      string `json:"Ticker" binding:"required"`
	Target_from string `json:"Target_from" binding:"required"`
	Target_to   string `json:"Target_to" binding:"required"`
	Company     string `json:"Company" binding:"required"`
	Action      string `json:"Action" binding:"required"`
	Brokerage   string `json:"Brokerage" binding:"required"`
	Rating_from string `json:"Rating_from" binding:"required"`
	Rating_to   string `json:"Rating_to" binding:"required"`
	Time        string `json:"Time" binding:"required"`
}
