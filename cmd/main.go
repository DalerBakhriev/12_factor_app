package main

import (
	"net/http"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Info("Hello, World!!!")

	port := os.Getenv("PORT")
	if len(port) == 0 {
		logrus.Fatal("Port was not set")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!!!"))
	})
	http.ListenAndServe(":"+port, nil)
}
