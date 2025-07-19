package perf

import (
	"testing"
	"time"
)

func TestAdjustOAT(t *testing.T) {
	got := AdjustOAT(25, time.Date(2023, time.January, 1, 12, 0, 0, 0, time.UTC))
	if got != 32 {
		t.Fatalf("expected 32 got %v", got)
	}
}

func seedPassengers() []string {
	return []string{"Anne", "Bob"}
}
