package repl

import (
	"fmt"
	"testing"
)

var historyEntryCmd = HistoryEntry{
	Command: Command{
		Name:        "test",
		Description: "test command",
		Handler: func(cfg Config, args []string) {
			fmt.Println("test")
		},
	},
	Args: []string{"test"},
}

func TestHistory_Add(t *testing.T) {
	r := newREPL()
	history := r.Config.History
	history.MaxSize = 3

	history.Add(historyEntryCmd)

	if len(*history.Entries) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(*history.Entries))
	}
	history.Add(historyEntryCmd)
	if len(*history.Entries) != 2 {
		t.Errorf("Expected 2 entries, got %d", len(*history.Entries))
	}
	history.Add(historyEntryCmd)
	if len(*history.Entries) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(*history.Entries))
	}
	history.Add(historyEntryCmd)
	if len(*history.Entries) != 3 {
		t.Errorf("Expected 3 entries, got %d", len(*history.Entries))
	}
}

func TestHistoryList(t *testing.T) {
	r := newREPL()
	history := r.Config.History
	entries := history.List()
	if len(entries) != 0 {
		t.Errorf("Expected 0 entries, got %d", len(entries))
	}
	history.Add(historyEntryCmd)
	entries = history.List()
	if len(entries) != 1 {
		t.Errorf("Expected 1 entry, got %d", len(entries))
	}

}
