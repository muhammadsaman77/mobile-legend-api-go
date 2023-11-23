package entities

type User struct {
	Id           int
	Name         string
	Email        string
	Password     string
	Gender       string
	VictoryCount int
	DefeatCount  int
	Winrate      int
}