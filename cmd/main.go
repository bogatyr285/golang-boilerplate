package main

import (
	"context"
	"log"

	"github.com/bogatyr285/golang-boilerplate/cmd/commands"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// no root cmd for now, but could be added
	cmd := commands.NewServeCmd()

	if err := cmd.ExecuteContext(ctx); err != nil {
		log.Fatalf("Something went wrong: %v", err)
	}
}
