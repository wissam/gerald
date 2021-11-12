package service

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/wissam/gerald/internal/commands"

	"github.com/gempir/go-twitch-irc/v2"
)

func Run() {
	client := twitch.NewClient(os.Getenv("NICKNAME"), os.Getenv("OAUTH"))
	client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		if strings.HasPrefix(message.Message, "!") {
			CommandsParser(message)
		} else {
			MessageParser(message)
		}
	})

	client.Join("Kodder")
	cerr := client.Connect()
	if cerr != nil {
		log.Fatal("Could not connect to twitch!")
		log.Fatal(cerr)
		panic(cerr)
	}
}
func CommandsParser(message twitch.PrivateMessage) {
	//The map should be easier to edit, maybe json? yaml? will figure it out.
	m := map[string]func(twitch.PrivateMessage) string{
		"hi":        commands.Hi,
		"hello":     commands.Hello,
		"favourite": commands.Favourite,
	}
	firstword := strings.Fields(message.Message)[0][1:]
	if val, exists := m[firstword]; exists {
		fmt.Println(val(message))
	}
}

func MessageParser(message twitch.PrivateMessage) {
	fmt.Println("Parse Message")
}
