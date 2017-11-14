package model

type TicTac struct {
	Plateau [][] string `json:"conditionals" sql:"type:json"`
	Turn    int
	Player1 string
	Player2 string
}
