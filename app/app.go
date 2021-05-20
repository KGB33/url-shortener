package app

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var ctx = context.Background()

func Entrypoint() {
	rdb, err := NewDBClient("localhost:6379", "", 0)
	env := server{rdb, mux.NewRouter()}
	if err != nil {
		log.Fatal(err)
	}
	env.run()
}

func NewDBClient(addr string, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
	if err := client.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return client, nil

}
