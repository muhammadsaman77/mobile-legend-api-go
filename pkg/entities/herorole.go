package entities

type HeroRole struct {
	Id     int
	HeroId int `json:"hero_id"`
	RoleId int `json:"role_id"`
}