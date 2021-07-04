package repeat

import (
	"context"
	"sync"
	"time"
)

// Start starts periodic calls in a goroutine.
// Returns function to stops that process.
func Start(period time.Duration, fn func(ctx context.Context)) (stop func()) {
	ctx, ctxCancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		run(ctx, period, fn)
	}()
	return func() {
		ctxCancel()
		wg.Wait()
	}
}

// StartWithContext starts periodic calls in a goroutine.
// That process is stopped when ctx is done.
func StartWithContext(ctx context.Context, period time.Duration, fn func(ctx context.Context)) {
	go run(ctx, period, fn)
}

func run(ctx context.Context, period time.Duration, fn func(ctx context.Context)) {
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
