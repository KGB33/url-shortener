package app

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

var ctx = context.Background()

type Server struct {
	DB     *redis.Client
	Router *mux.Router
}

func (s *Server) Run(port string) {
	fmt.Println("Starting Server...")
	defer s.DB.ShutdownSave(ctx)
	log.Fatal(http.ListenAndServe(port, s.Router))
}

func NewServer(db_addr string, db_pass string, db_id int) *Server {
	rdb, err := NewDBClient(db_addr, db_pass, db_id)
	env := Server{rdb, mux.NewRouter()}
	if err != nil {
		log.Fatal(err)
	}
	env.initRoutes()
	return &env
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
