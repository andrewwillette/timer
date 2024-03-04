package main

import (
	"testing"
	"time"
)

func TestGetDuration(t *testing.T) {
	t.Run("getDuration", func(t *testing.T) {
		got := getDuration("5s")
		want := 5 * time.Second
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
