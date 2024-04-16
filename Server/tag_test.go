package Server

import (
	pb "Grpclearn/Proto"
	"context"
	"log"
	"testing"
)

func TestTagserver_GetTageList(t *testing.T) {
	req := &pb.GetTaglistReques{Name: "weatherInfo"}
	Tagser := NewTageServer()
	reply, err := Tagser.GetTageList(context.Background(), req)
	if err != nil {
		log.Printf(err.Error())
	}
	log.Println(reply)
	log.Println(reply.GetLives())
}
