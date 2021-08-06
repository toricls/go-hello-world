package main

import (
	"fmt"
	"net/http"
	"os"

	"go.uber.org/zap"
)

func getMessage(defaultValue string) string {
	if value, ok := os.LookupEnv("MESSAGE"); ok {
		return value
	}
	return defaultValue
}

func getPortNumber(defaultValue string) string {

	if value, ok := os.LookupEnv("PORT_NUMBER"); ok {
		return value
	}
	return defaultValue
}

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	logger.Info("starting up ... 🤘")

	msg := getMessage("Hello, World!")
	listenAddr := fmt.Sprintf(":%s", getPortNumber("80"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, msg)
		logger.Info("I'm about to say something for you",
			zap.String("method", r.Method),
			zap.String("path", r.URL.Path),
			zap.String("ua", r.UserAgent()),
		)
	})
	if err := http.ListenAndServe(listenAddr, nil); err != nil {
		logger.Fatal("failed to listen and serve",
			zap.Error(err),
		)
	}
}
