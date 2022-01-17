package tirc

import (
	"log"

	"github.com/gempir/go-twitch-irc/v2"
	"github.com/wissam/gerald/internal/models"
)

var client twitch.Client
var channels []models.Channel

func (c *client) Run() {
	log.Println("Connecting")
}

func (c *client) AddChannel(channel models.Channel) {
	log.Println("Adding Channel")
}

func (c *client) DeleteChannel(channel models.Channel) {
	log.Println("Removing Channel")
}

func (c *client) JoinChannel(channel models.Channel) {
	log.Println("Joining Channel")
}

func (c *client) LeaveChannel(channel models.Channel) {
	log.Println("Leaving Channel")
}
