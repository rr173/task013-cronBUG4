package cron

import (
	"testing"
	"time"
)

// TestNextNCapExactly100 verifies that NextN with count > 100 is capped at
// exactly 100, not 99. This catches an off-by-one in the upper bound clamp.
func TestNextNCapExactly100(t *testing.T) {
	s, _ := Parse("* * * * *")
	from := time.Date(2026, 1, 15, 10, 0, 0, 0, time.UTC)
	got, err := s.NextN(from, 200)
	if err != nil {
		t.Fatalf("NextN: %v", err)
	}
	if len(got) != 100 {
		t.Errorf("NextN(200) returned %d items, want exactly 100", len(got))
	}
}

// TestNextNCapBoundary verifies the boundary: count=100 should return 100.
func TestNextNCapBoundary(t *testing.T) {
	s, _ := Parse("* * * * *")
	from := time.Date(2026, 1, 15, 10, 0, 0, 0, time.UTC)
	got, err := s.NextN(from, 100)
	if err != nil {
		t.Fatalf("NextN: %v", err)
	}
	if len(got) != 100 {
		t.Errorf("NextN(100) returned %d items, want exactly 100", len(got))
	}
}

// TestNextNCapAt101 verifies count=101 is capped to 100.
func TestNextNCapAt101(t *testing.T) {
	s, _ := Parse("* * * * *")
	from := time.Date(2026, 1, 15, 10, 0, 0, 0, time.UTC)
	got, err := s.NextN(from, 101)
	if err != nil {
		t.Fatalf("NextN: %v", err)
	}
	if len(got) != 100 {
		t.Errorf("NextN(101) returned %d items, want exactly 100", len(got))
	}
}
