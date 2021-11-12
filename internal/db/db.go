package db

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Initialisiation should be on demand not everytime, I think?
// RoR does that, I wonder if gorm is different.
// If yes, should implement some sort of basic cli? huhmmm

//Basic CRUD + migration + initialisation? do I need to initialise if I have
//migrate?  probably not.
// Meh, keeping it for now.
// I wonder if I will get a dmca strike on my first stream...weeeee
// Yep initialise is for connection...good stuff
type User struct {
	gorm.Model
	twitch_id      int
	display_name   string
	name           string
	bio            string
	email_verified bool
	created_at     *time.Time
	role_channel   string
	partnered      bool
	user_type      string
	updated_at     *time.Time
	Emotes         []*Emote `gorm:"many2many:user_emotes;"`
}

//how often should I check for updates? event based?

type Channel struct {
	gorm.Model
	broadcaster_id       int
	broadcaster_login    string
	broadcaster_name     string
	broadcaster_language string
	game_id              int
	game_name            string
	title                string
	delay                int
}

type Emote struct {
	gorm.Model
	emote_type string //channel or global
	Users      []*User `gorm:"many2many:user_emotes;"`
}

// need to link them together with foreign keys, need to brush up on db
// normalisation too... forget all this.

var DBCon *gorm.DB //best practice regarding global variables in golang/gorm? is this normal?

func Initialise() {
	var err error
	fmt.Println("Begin")
	DBCon, err = gorm.Open(sqlite.Open("gerald.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
}

func Migrate() {
	fmt.Println("Migrate")
	//add structs to migrate as they go.
	// need to check migrate vs automigrate in gorm. there should be a
	// difference.
	//answer: automigrate just detects differences and applies itself, there
	//are some limitations like won't delete and won't update columns. not a
	//big issue at the moment with sqlite, but have to learn this when I move
	//to mysql(or whatever db I decide) and move to production.
	DBCon.AutoMigrate(&User{})
	DBCon.AutoMigrate(&Channel{})
	DBCon.AutoMigrate(&Emote{})
	

func New_User() {
	fmt.Println("Create new user")
}

func Update_User() {
	fmt.Println("Update user")
}

func Get_User() {
	fmt.Println("Get user")
}

func Delete_User() {
	fmt.Println("Delete user")
}
