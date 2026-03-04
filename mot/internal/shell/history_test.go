package shell

import "testing"

func TestHistoryAppend(t *testing.T) {
	s := &Shell{}

	s.history = append(s.history, "ls")
	s.history = append(s.history, "pwd")

	if len(s.history) != 2 {
		t.Fatalf("expected 2 entries, got %d", len(s.history))
	}

	if s.history[0] != "ls" {
		t.Fatalf("expected 'ls', got %s", s.history[0])
	}
}
