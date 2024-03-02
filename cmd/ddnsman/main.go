package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"os/signal"

	"github.com/leonidboykov/ddnsman"
)

func main() {
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, nil)))

	config, err := ddnsman.LoadConfiguration()
	if err != nil {
		log.Fatalln("read configuration:", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, os.Kill)
	defer cancel()

	u, _ := ddnsman.New(config)
	if err := u.Start(ctx); err != nil {
		log.Fatalln(err)
	}
}
