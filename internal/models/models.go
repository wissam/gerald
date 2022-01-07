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

// still don't like the name...
type Rcred struct {
	Address  string
	Password string //plain text string for now, encrypt before production release
	Database int
	channel  string
}

// todo tag with json and gorm (I think)
