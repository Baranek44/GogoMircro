package main

import (
	"context"
	"log"
	"logger-service/data"
	"time"
)

type RPCServer struct {
}

type RPCPayload struct {
	Name string
	Data string
}

func (s *RPCServer) LogInfo(payload RPCPayload, resp *string) error {
	coll := client.Database("logs").Collection("logs")
	_, err := coll.InsertOne(context.TODO(), data.LogEntry{
		Name:      payload.Name,
		Data:      payload.Data,
		CreatedAt: time.Now(),
	})

	if err != nil {
		log.Println("Error writing to mongo", err)
		return err
	}

	*resp = "Processed payload via RPC" + payload.Name
	return nil
}
