package bot

import (
	"log"
	"os"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/wissam/gerald/internal/models"
	"github.com/wissam/gerald/internal/vislog" //I am very confused...
)

var client *twitch.Client

func Run() {
	channels := []models.Channel{{Name: "Kodder", ID: "101185038"}}
	client = twitch.NewClient(os.Getenv("NICKNAME"), os.Getenv("OAUTH"))
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		NickBulbColour(message.User.Color)
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

func NickBulbColour(color string) {
	blb := vislog.NewBulb("d073d567639b")
	blb.HEX(color)
}
