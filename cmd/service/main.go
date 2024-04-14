package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"

	"service-dialog/cmd"
	"service-dialog/internal/cache"
	dto "service-dialog/internal/generated"
	"service-dialog/internal/rpc/handlers/create"
	"service-dialog/internal/rpc/handlers/get"
	"service-dialog/internal/server"
	"service-dialog/internal/storage"
	"service-dialog/internal/usecase/message_manager"
)

func main() {
	ctx := context.Background()

	storage := storage.New(cmd.MustInitPostgresql())
	dialogsCache := cache.Must(ctx, storage)

	messageManager := message_manager.New(storage, dialogsCache)

	createHandler := create.New(messageManager)
	getHandler := get.New(messageManager)

	service := grpc.NewServer()
	srv := server.New(
		createHandler,
		getHandler,
	)
	dto.RegisterServiceDialogServer(service, srv)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatal("Unable to create grpc listener:", err)
	}

	if err = service.Serve(listener); err != nil {
		log.Fatal("Unable to start server:", err)
	}
}
