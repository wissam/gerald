package db

import (
	//"fmt"
	"log"

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
// need to find some defaults for now, and how gorm handles them
//type User struct {
//	gorm.Model
//	TwitchID string   `gorm:"unique"`
//	Emotes   []*Emote `gorm:"many2many:user_emotes;"` //no duplicate users!
//}

//Seems the library is incomplete... maybe? the api clearly declares those
//variables but the lib doesn't have them....okie...digging more...
// yep library is incomplete... two choices now. make a layer on top to get
// those details myself or help the creator of lib? but i am still too new with
// golang to do that... should put it as a todo.
//how often should I check for updates? event based?

// Scrap all this, this is irc! find another lib to deal with the twitch api on
// application level!!! gah!  google google google!

// I wonder if I should have many to many with all tables.
// a USER can be in multiple channels, and use multiple EMOTES
// a CHANNEL can have multiple USERS and multiple EMOTES
// An EMOTE can be used by multiple USERS and in multiple CHANNELS... as long
// as the user is a sub for the channel that the emote belongs too... okie this
// is getting a bit more complicated for someone who has not touched databases
// in 15 years... I do not remember anything...
// okay... I will make it easy on myself for now or else I will not accomplish
// anything today.. forget "channels"... let's assume the bot goes to only 1
// channel for now. (it does heh...it still hardcoded to here)

//type Channel struct {
//	gorm.Model
//	broadcaster_id       int
//	broadcaster_login    string
//	broadcaster_name     string
//	broadcaster_language string
//	game_id              int
//	game_name            string
//	title                string
//	delay                int
//}
// this is wrong, emotes should be related to a join between user and channel.
// preferred emotes are dependent on where they are being used...
// I don't remember dbs at all...
// I miss listening to this song...weeee
//type Emote struct {
//	gorm.Model
//	Name  string
//	EId   string
//	Count int
//	Users []*User `gorm:"many2many:user_emotes;"`
//}

// need to link them together with foreign keys, need to brush up on db
// normalisation too... forget all this.

// At one point I have to make a db "object"
// I guess now it is the time to create the object!
//damn I am so sleepy... this flu is not gone yet...
//var DBCon *gorm.DB //best practice regarding global variables in golang/gorm? is this normal?
//reading a bit on how to implrement this...brb
//--------------------------------------------------------------------
// Temporary type
type Counter struct {
	gorm.Model
	UserId string
	RoomId string
	Emote  string
	Count  int
}

var DBCon *gorm.DB

func Connect() {
	var err error
	DBCon, err = gorm.Open(sqlite.Open("labase.db"), &gorm.Config{})
	if err != nil {
		log.Printf("Could not connect to database: %s\n", err)
		panic("Failed to connect to database")
	} else {
		log.Println("Connected to database")
	}
}

func Migrate() {
	log.Println("Migrating Database")
	//add structs to migrate as they go.
	// need to check migrate vs automigrate in gorm. there should be a
	// difference.
	//answer: automigrate just detects differences and applies itself, there
	//are some limitations like won't delete and won't update columns. not a
	//big issue at the moment with sqlite, but have to learn this when I move
	//to mysql(or whatever db I decide) and move to production.
	//DBCon.AutoMigrate(&User{})
	//DBCon.AutoMigrate(&Channel{})
	//DBCon.AutoMigrate(&Emote{})
	DBCon.AutoMigrate(&Counter{})
}

// do I need to store what I can get from twitch? or is it a wrong pattern? if
// I store it, then I would have to have a mechanism to update it all the
// time.. I wonder if this is stupid. maybe I should save the id and make the
// api work for me for display.
// yeah better just store id, in case there is a change of nickname,displayname
// or color, there is no point of keeping it.
//func New_User(twitch_id string) (ID uint) {
//	var user User
//	result := DBCon.Where("twitch_id = ?", twitch_id).First(&user)
//	if result.Error != nil {
//
//		user = User{TwitchID: twitch_id}
//		DBCon.Create(&user)
//	}
//	fmt.Println(user.ID) //?????
//	return user.ID
//}

// I think the many to mnay relationship is wrong... I need a
// user , emote , count table... I wonder how to create that...
// and this would solve my problem for channel... I think...
// so user, channel, emote, count ...I think... is this a normal table?
// I really need to learn proper db shit...
//func Get_User() {
//	fmt.Println("Get user")
//}

// why would I delete a user? meh... remove...
//func Delete_User() {
//	fmt.Println("Delete user")
//}
//--------------------------------------------------------------------
// Above will be reactivated once I have a functional structure
// Below will removed later on

// Insert into a flat table... nothing fancy
// user id, room id, emote , count

func EmoteCountInsert(userid string, roomid string, emote string, count int) {
	var counter Counter
	result := DBCon.Where("user_id = ? AND room_id = ? AND emote = ?", userid, roomid, emote).First(&counter)
	if result.Error != nil {
		counter = Counter{UserId: userid, RoomId: roomid, Emote: emote, Count: count}
		DBCon.Create(&counter)
	} else {
		newcount := counter.Count + count
		DBCon.Model(&counter).Update("count", newcount) // can I do arithmetrics here?
	}
}

func EmoteCountGetter(userid string, roomid string) {
	//var counter Counter
	var counters []Counter

	result := DBCon.Where("user_id = ? AND room_id = ?", userid, roomid).Find(&counters)

	log.Printf("Rows Affected %d\n", result.RowsAffected)
}
