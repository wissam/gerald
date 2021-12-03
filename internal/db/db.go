package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Counter struct {
	gorm.Model
	UserId string
	RoomId string
	Emote  string
	Count  int
}

type DB struct {
	db *gorm.DB
}

func (d *DB) Connect() {
	var err error
	d.db, err = gorm.Open(sqlite.Open("labase.db"), &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)})
	if err != nil {
		log.Printf("Could not connect to database: %s\n", err)
		panic("Failed to connect to database")
	} else {
		log.Println("Connected to database")
	}
	log.Println(d.db)
}

func (d *DB) Migrate() {
	log.Println("Migrating Database")
	d.db.AutoMigrate(&Counter{})
}

func (d *DB) EmoteCountInsert(userid string, roomid string, emote string, count int) {
	var counter Counter
	result := d.db.Where("user_id = ? AND room_id = ? AND emote = ?", userid, roomid, emote).First(&counter)
	if result.Error != nil {
		counter = Counter{UserId: userid, RoomId: roomid, Emote: emote, Count: count}
		d.db.Create(&counter)
	} else {
		newcount := counter.Count + count
		d.db.Model(&counter).Update("count", newcount) // can I do arithmetrics here?
	}
}

func (d *DB) EmoteCountGetter(userid string, roomid string) {
	//var counter Counter
	var counters []Counter

	result := d.db.Where("user_id = ? AND room_id = ?", userid, roomid).Find(&counters)

	log.Printf("Rows Affected %d\n", result.RowsAffected)
}
