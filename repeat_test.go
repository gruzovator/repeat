package repeat_test

import (
	"context"
	"testing"
	"time"

	"github.com/gruzovator/repeat"
)

func TestStart(t *testing.T) {
	doneCh := make(chan struct{})
	wantCallsNum := 10
	callsNum := 0

	stopFn := repeat.Start(time.Millisecond, func(ctx context.Context) {
		callsNum++
		if callsNum == wantCallsNum {
			close(doneCh)
		}
	})
	select {
	case <-doneCh:
	case <-time.After(time.Duration(wantCallsNum)*time.Millisecond + 10*time.Millisecond):
	}
	stopFn()
	time.Sleep(10 * time.Millisecond)

	if callsNum != 10 {
		t.Fatalf("calls num: want: %d, got:%d", wantCallsNum, callsNum)
	}
}

func TestStartWithContext(t *testing.T) {
	doneCh := make(chan struct{})
	wantCallsNum := 10
	callsNum := 0
	ctx, ctxCancel := context.WithCancel(context.Background())

	repeat.StartWithContext(ctx, time.Millisecond, func(ctx context.Context) {
		callsNum++
		if callsNum == wantCallsNum {
			close(doneCh)
		}
	})
	select {
	case <-doneCh:
	case <-time.After(time.Duration(wantCallsNum)*time.Millisecond + 10*time.Millisecond):
	}
	ctxCancel()
	time.Sleep(10 * time.Millisecond)

	if callsNum != 10 {
		t.Fatalf("calls num: want: %d, got:%d", wantCallsNum, callsNum)
	}
}
