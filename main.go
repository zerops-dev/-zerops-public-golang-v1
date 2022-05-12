package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	Version string
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	signals := make(chan os.Signal)
	signal.Notify(signals, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		<-signals
		log.Println("calling cancel")
		cancel()
	}()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		log.Println("request accepted")
		fmt.Fprint(writer, "version: " + Version)
	})

	go func() {
		log.Println("listening on :8080")
		if err := http.ListenAndServe(":8080", nil); err != nil {
			fmt.Println(err.Error())
		}
	}()

	<-ctx.Done()
	log.Println("ctx done")
}
