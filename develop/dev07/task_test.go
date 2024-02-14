package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	c1 := make(chan interface{})
	c2 := make(chan interface{})
	c3 := make(chan interface{})

	go func() {
		time.Sleep(2 * time.Second)
		close(c1)
	}()
	go func() {
		time.Sleep(3 * time.Second)
		close(c2)
	}()
	go func() {
		time.Sleep(1 * time.Second)
		close(c3)
	}()

	done := or(c1, c2, c3)

	start := time.Now()
	<-done
	duration := time.Since(start)

	if duration < 1*time.Second || duration > 3*time.Second {
		t.Errorf("Unexpected duration: %v", duration)
	}
}

func TestSig(t *testing.T) {
	start := time.Now()
	<-sig(2 * time.Second)
	duration := time.Since(start)

	if duration < 2*time.Second {
		t.Errorf("Unexpected duration: %v", duration)
	}
}
