package repeat_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"github.com/gruzovator/repeat"
)

func TestRun(t *testing.T) {
	t.Run("does periodic calls", func(t *testing.T) {
		ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancelCtx()

		callsCounter := 0
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			repeat.Run(ctx, time.Millisecond, func(ctx context.Context) {
				callsCounter++
			})
		}()
		wg.Wait()

		if callsCounter <= 1 {
			t.Fatalf("no periodic calls: calls number: %d", callsCounter)
		}
	})

	t.Run("fn context is cancelled when Run context is cancelled", func(t *testing.T) {
		ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer cancelCtx()

		doneCh := make(chan struct{})
		go func() {
			defer close(doneCh)
			repeat.Run(ctx, time.Millisecond, func(ctx context.Context) {
				<-ctx.Done()
			})
		}()

		select {
		case <-doneCh:
		case <-time.After(100 * time.Millisecond):
			t.Fatalf("fn context is not cancelled")
		}
	})
}
