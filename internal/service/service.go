package service

import (
	"log"
	"os"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/wissam/gerald/internal/commands"
	"github.com/wissam/gerald/internal/db"
	"github.com/wissam/gerald/internal/tapi"
)

//temporary struct channel until I figure out the db
// got pomodoro off stream...good enough for now...
type Channel struct {
	Name   string
	ID     string
	Emotes []string //array of strings for now, but it will be of structs prob
}

type Emote struct {
	ID    string
	Name  string
	InUse bool
}

var dbi db.DB
var client *twitch.Client

//yep good enough for now...

func Listen() {
	log.Println("Listening")
}

func Run() {
	//I forgot what I was doing...hmmm, let's see...
	dbi = db.DB{}
	dbi.Connect()
	dbi.Migrate()
	//this whole design is shitty...
	// unsure what's the best way to deal wiht it...

	// So I can assume that the id will come from the database somehow?
	// the "how" will be determined later.
	// I need a join for all those ids , a emotes retrievals for all those ids.
	// then a check
	// I wonder, does twitch send an event trigger upon channel starting
	// broadcast?
	// Thinking...not typing anything...
	//-------------
	//-------------tapi.GetUser("kodder") //forgot this, I think the Getuser is gone...
	//-------------
	channels := []Channel{{Name: "MadMistro", ID: "38429111"}, {Name: "Kodder", ID: "101185038"}}
	for i := 0; i < len(channels); i++ {
		channels[i].Emotes = tapi.GetAllEmotes(channels[i].ID)
		log.Printf("len inside first loop %d\n", len(channels[i].Emotes))
	}
	for _, channel := range channels {

		log.Printf("len second loop %d\n", len(channel.Emotes))
		for _, emote := range channel.Emotes {

			log.Println(emote)
		}
	}
	client = twitch.NewClient(os.Getenv("NICKNAME"), os.Getenv("OAUTH"))
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if strings.HasPrefix(message.Message, "!") {
			CommandsParser(message)
		} else {
			MessageParser(message)
		}
	})
	for _, channel := range channels {
		client.Join(channel.Name)
	}
	cerr := client.Connect() //this blocks anything under it, doesn't reach the if/else
	if cerr != nil {
		log.Fatal("Could not connect to twitch!")
		log.Fatal(cerr)
		panic(cerr)
	}
}
func CommandsParser(message twitch.PrivateMessage) {

	m := map[string]func(twitch.PrivateMessage, db.DB) string{
		"hi":        commands.Hi,
		"hello":     commands.Hello,
		"favourite": commands.Favourite,
		"emotes":    commands.Emotes,
	}
	firstword := strings.Fields(message.Message)[0][1:]
	if val, exists := m[firstword]; exists {
		client.Say(message.Channel, val(message, dbi))
	}
}

//I need to add emotes only related to the room/channel that is currently being
//used , or you would have to give the bot subscription in every channel where
//any emote has been used ever... gosh that's a lot of money for nothing...
// So design change, maybe I should only have a counter with the emotes only?
// Let's see how to get the ownership of emote before thinking about anything
// else...
func MessageParser(message twitch.PrivateMessage) {
	if message.Emotes != nil {
		for _, e := range message.Emotes {
			dbi.EmoteCountInsert(message.User.ID, message.RoomID, e.ID, e.Count)
		}
	}
}
