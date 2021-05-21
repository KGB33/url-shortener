package app

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var ctx = context.Background()

type server struct {
	db     *redis.Client
	router *mux.Router
}

func (s *server) Run(port string) {
	defer s.db.ShutdownSave(ctx)
	s.routes(port)
}

func NewServer(db_addr string, db_pass string, db_id int) server {
	rdb, err := NewDBClient(db_addr, db_pass, db_id)
	env := server{rdb, mux.NewRouter()}
	if err != nil {
		log.Fatal(err)
	}
	return env
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
