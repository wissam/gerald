package commands

import (
	"fmt"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/wissam/gerald/internal/db"
)

//interesting, it doesn't seem to complain if there is an unused argument...
// is it a bug? or something i still need to check??

func Hi(message twitch.PrivateMessage, dbi db.DB) string {
	return fmt.Sprintf("Hi %s\n", message.User.DisplayName)
}

func Hello(message twitch.PrivateMessage, dbi db.DB) string {
	return "hello"
}

func Favourite(message twitch.PrivateMessage, dbi db.DB) string {
	return "fav"
}

func Emotes(message twitch.PrivateMessage, dbi db.DB) string {
	emotes := dbi.EmoteCountGetter(message.User.ID, message.RoomID)
	return fmt.Sprintf("%s these are all the emotes you have used %s", message.User.DisplayName, emotes)
}

//I need to access the db, what's the best practice here? connection again or
// pass the pointer along side the entire db package? no clue...
// time to investigate...:
