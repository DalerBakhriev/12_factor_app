package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func main() {

	log := logrus.New()
	log.SetOutput(os.Stdout)
	log.Info("Starting the application")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		log.Fatal("Port was not set")
	}

	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, Ari!!!"))
	})

	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: router,
	}

	go server.ListenAndServe()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	<-interrupt

	log.Info("Stopping the application")

	timeout, cancelFunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelFunc()
	err := server.Shutdown(timeout)
	if err != nil {
		log.Error("Error while shutting down: %v", err)
	}
	log.Info("The application was stopped")
}
