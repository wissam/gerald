package bot

import (
	"context"
	"log"
	"os"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/go-redis/redis/v8"
	"github.com/wissam/gerald/internal/commands"
	"github.com/wissam/gerald/internal/db"
	"github.com/wissam/gerald/internal/models"
	"github.com/wissam/gerald/internal/tapi"
)

//********************************************************************************************************
// Migrate all DB stuff to db module
// Migrate all irc stuff to irc module
// Migrate all redis stuff to redis module
//temporary struct channel until I figure out the db
// I should probably have models first... then build on top bi by bit to tame
// the chaos...
var dbi db.DB
var client *twitch.Client

//yep good enough for now...

func Listen() {
	log.Println("Listening")
}

func Run() {
	dbi = db.DB{}
	dbi.Connect()
	dbi.Migrate()
	//temporary redis connect , and sub
	// DO NOT KEEP!!!!
	// **********************************************************************
	var ctx = context.Background()
	//rmsg := make(chan redis.Message)
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	newchan := make(chan string)
	go func() {
		subscriber := rdb.Subscribe(ctx, "one")
		for {
			log.Println("looping")
			rmsg := subscriber.Channel()
			select {
			case rchan, ok := <-rmsg:
				if ok {
					log.Printf("Got a redis payload with %s\n", rchan.Payload)
					newchan <- rchan.Payload
					log.Println("unblocked payload and went to newchan")
				}
			}
		}
	}()

	channels := []models.Channel{{Name: "MadMistro", ID: "38429111"}, {Name: "Kodder", ID: "101185038"}}
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
	// lifewithbugs is a hero!
	// Learn the proper convention for "listen loop" or "waiting loop"
	go func() {
		for {
			AddChannel(newchan)
		}
	}()
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
	// Check if you can put this in a goroutine
	cerr := client.Connect() //this blocks anything under it, doesn't reach the if/else
	if cerr != nil {
		log.Fatal("Could not connect to twitch!")
		log.Fatal(cerr)
		panic(cerr)
	}
}

//I am not sure how design this...
func AddChannel(nchan chan string) {
	log.Println("pre channel assignment to see if this is getting blocked")
	//this needs a select with a cases. including "done".
	newchan := <-nchan
	log.Printf("Attempting to join %s\n", newchan)
	client.Join(newchan)
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
	log.Println(message.Channel) //this is magic, it solves all problems on earth...
}
