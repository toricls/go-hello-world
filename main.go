package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/aws/aws-lambda-go/lambda"

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

func getSleepMillisec(defaultValue int) int {
	if valueStr, ok := os.LookupEnv("SLEEP_MILLISEC"); ok {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return defaultValue
}

func isLambdaEnvinronment() bool {
	// https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-runtime
	_, present := os.LookupEnv("_HANDLER")
	return present
}

func lambdaHandler(ctx context.Context) (string, error) {
	if sleepMsec > 0 {
		time.Sleep(time.Duration(sleepMsec) * time.Millisecond)
	}
	return msg, nil
}

var msg = ""
var listenAddr = "80"
var sleepMsec = 0

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	msg = getMessage("Hello, World!")
	listenAddr = fmt.Sprintf(":%s", getPortNumber("80"))
	sleepMsec = getSleepMillisec(0)

	if isLambdaEnvinronment() {
		// Run as Lambda function
		lambda.Start(lambdaHandler)
	} else {
		// Run as normal http server
		logger.Info("starting up ... 🤘",
			zap.String("message", msg),
			zap.String("listen-address", listenAddr),
			zap.Int("sleep-millisec", sleepMsec),
		)

		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if sleepMsec > 0 {
				time.Sleep(time.Duration(sleepMsec) * time.Millisecond)
			}
			fmt.Fprint(w, msg)
			logger.Info("I just said something for you",
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
}
