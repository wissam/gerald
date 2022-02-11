package models

type Channel struct {
	Name   string
	ID     string
	Emotes []Emote
}

type Emote struct {
	ID    string
	Name  string
	InUse bool
}

// Write CRUD , do I use gorm or not?
