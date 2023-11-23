package entities

type Hero struct {
	Id              int     `json:"id"`
	Name            string  `json:"name"`
	BasicHp         float32 `json:"basic_hp"`
	PhysicalAttack  float32 `json:"physical_attack"`
	MagicalAttack   float32 `json:"magical_attack"`
	PhysicalDefense float32 `json:"physical_defense"`
	MagicalDefense  float32 `json:"magical_defense"`
	Roles           string  `json:"roles"`
}
