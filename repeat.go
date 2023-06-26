package repeat

import (
	"context"
	"sync"
	"time"
)

// Run runs periodic calls of a function fn.
// It stops when the ctx is done.
func Run(ctx context.Context, period time.Duration, fn func(ctx context.Context)) {
	ticker := time.NewTicker(period)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			fn(ctx)
		case <-ctx.Done():
			return
		}
	}
}

// Start starts periodic calls in a goroutine.
// Returns a function to stop that process.
func Start(period time.Duration, fn func(ctx context.Context)) (stop func()) {
	ctx, ctxCancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		Run(ctx, period, fn)
	}()
	return func() {
		ctxCancel()
		wg.Wait()
	}
}
