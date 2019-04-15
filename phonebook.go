package phonebook

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

// Phonebook represents a phonebook server
type Phonebook struct {
}

// New creates a new phonebook server
func New() *Phonebook {
	return &Phonebook{}
}

// Run runs the phonebook server, which starts an http listener. Blocks execution.
func (p *Phonebook) Run(ctx context.Context) error {
	g := gin.New()
	exited := make(chan struct{})
	var err error
	go func() {
		err = g.Run("localhost:8080")
		exited <- struct{}{}
	}()
	select {
	case <-exited:
		log.Println("Listen exited with error:", err)
		return err
	case <-ctx.Done():
		log.Println("Cancelled")
		return nil
	}
}
