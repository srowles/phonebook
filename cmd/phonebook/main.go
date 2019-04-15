package main

import (
	"context"

	"github.com/srowles/phonebook"
)

func main() {
	server := phonebook.New()
	server.Run(context.Background())
}
