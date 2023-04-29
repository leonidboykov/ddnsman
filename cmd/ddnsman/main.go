package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/leonidboykov/ddnsman"
)

func main() {
	config, err := ddnsman.LoadConfiguration("config.json") // TODO: Make it customizable.
	if err != nil {
		log.Fatalln("unable to read configuration:", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	u, _ := ddnsman.New(config)
	if err := u.Start(ctx); err != nil {
		log.Fatalln(err)
	}
}
