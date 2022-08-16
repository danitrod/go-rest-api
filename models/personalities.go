package models

type Personality struct {
	Name  string `json:"name"`
	Story string `json:"story"`
	Id    int    `json:"id"`
}

var Personalities []Personality
