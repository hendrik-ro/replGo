package repl

func (h *History) Add(entry HistoryEntry) {
	// Adds a new entry to the history, truncating at max size.
	*h.Entries = append(*h.Entries, entry)
	if len(*h.Entries) > h.MaxSize {
		*h.Entries = (*h.Entries)[1:]
	}
}

func (h *History) List() []HistoryEntry {
	// Returns a list of all entries in the history.
	return *h.Entries
}
