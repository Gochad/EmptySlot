package server

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func ServerConfig(port string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    port,
		Handler: handler,
	}
}

func CloseServer(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	err := server.Shutdown(nil)
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

func NewServer(r http.Handler) {
	s := ServerConfig(":8080", r)

	RunServer(s)
	CloseServer(s)
}
