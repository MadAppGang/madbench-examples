package session

// Store is an in-memory session table. Unrelated to HandleLogin — it
// exists so the fixture has more than one package for the agent to
// skim past.
type Store struct {
	items map[string]string
}

// New returns an empty Store.
func New() *Store {
	return &Store{items: map[string]string{}}
}
