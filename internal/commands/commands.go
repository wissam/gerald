package commands

import (
	"github.com/gempir/go-twitch-irc/v2"
)

//note: user id while useful for the db lookup, won't be useful for direct
//interaction with the bot, better have a mechanism that separates what is
//needed from the user struct.
// In this case I need the displayname instead of Id!

func Hi(message twitch.PrivateMessage) string {
	return "hi"
}

func Hello(message twitch.PrivateMessage) string {
	return "hello"
}

func Favourite(message twitch.PrivateMessage) string {
	return "fav"
}

func Emotes(message twitch.PrivateMessage) string {
	return "Emotes"
}

//I need to access the db, what's the best practice here? connection again or
// pass the pointer along side the entire db package? no clue...
// time to investigate...:
