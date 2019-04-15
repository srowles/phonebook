package phonebook

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	p := New()
	require.NotNil(t, p, "Phonebook should be created")
}

func TestRun(t *testing.T) {
	p := New()
	require.NotNil(t, p, "Phonebook should be created")
	ctx, cancel := context.WithCancel(context.Background())
	exited := false
	go func() {
		err := p.Run(ctx)
		exited = true
		require.NoError(t, err, "should not error")
	}()

	time.Sleep(10 * time.Millisecond)
	require.False(t, exited, "should not have exited yet")
	cancel()
	time.Sleep(10 * time.Millisecond)

	require.True(t, exited, "running should have exited")
}
