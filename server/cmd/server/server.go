package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func Config(port string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    port,
		Handler: handler,
	}
}

func CloseServer(ctx context.Context, server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("Server closing error: %v", err)
	}
}

func RunServer(server *http.Server) {
	go func() {
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()
}

func NewServer(ctx context.Context, r http.Handler) {
	s := Config(":8080", r)

	RunServer(s)
	CloseServer(ctx, s)
}
