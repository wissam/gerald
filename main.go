package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gempir/go-twitch-irc/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {

	type User struct {
		gorm.Model
		Name string
	}
	db, err := gorm.Open(sqlite.Open("gerald.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&User{})
	db.Create(&User{Name: "Kodder"})
	var user User
	db.First(&user, 1)
	fmt.Printf("Name is %s\n", user.Name)

	client := twitch.NewClient(os.Getenv("NICKNAME"), os.Getenv("OAUTH"))

	client.OnPrivateMessage(func(message twitch.PrivateMessage) {

		if strings.HasPrefix(message.Message, "!") {
			BotCommands(message, client)
			//time.Sleep(10 * time.Second)
		} else {
			fmt.Println("the else")

			fmt.Printf("%s said: ", message.User.DisplayName)
			fmt.Println(message.Message)
			if message.Emotes != nil {
				fmt.Println("The message has the following emotes:")
				for _, e := range message.Emotes {
					fmt.Print(e.Name)
					fmt.Printf(" %d times \n", e.Count)
				}
			}
			fmt.Printf("The message type is %d ", message.Type)
		}
	})

	client.OnNoticeMessage(func(message twitch.NoticeMessage) {
		fmt.Printf("This is a notice message %s: ", message.Message)
	})

	client.OnUserNoticeMessage(func(message twitch.UserNoticeMessage) {
		fmt.Printf("This is a user notice message %s: ", message.Message)

	})
	client.OnRoomStateMessage(func(message twitch.RoomStateMessage) {
		for k, v := range message.Tags {
			fmt.Printf("%s   %s \n", k, v)
		}

		for k, v := range message.State {
			fmt.Printf("%s   %d \n", k, v)
		}
	})

	client.OnGlobalUserStateMessage(func(message twitch.GlobalUserStateMessage) {
		fmt.Printf("This is a global user state message %s\n", message.User.DisplayName)
	})
	client.Join("Kodder")

	ierr := client.Connect()
	if err != nil {
		panic(ierr)
	}
}

func BotCommands(message twitch.PrivateMessage, client *twitch.Client) {

}
