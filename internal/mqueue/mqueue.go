package mqueue

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
)

var rdb redis.Client //I wonder if this is the correct way to do it?
//googling
var ctx context.Context // hmm?...this is wrong, global and then getting passed...
// I read in the best practice page that this isn't great...so I guess I will
// initialise it in connect, then return a pointer ?
//need to pass context too
func (r *rdb) Connect() context.Context {
	log.Println("Connecting to Redis")
	r = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDISHOST"),
		Password: os.Getenv("REDISPASSWORD"), // no password set
		DB:       os.Getenv("REDISDB"),       // use default DB
	})
	return context.Background()
}

func (r *rdb) Listen(ctx context.Context) {
	log.Println("Listenning to Redis Channel")
}
