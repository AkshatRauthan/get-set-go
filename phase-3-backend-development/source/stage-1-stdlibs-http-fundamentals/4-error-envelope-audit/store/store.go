package store

import "sync"

type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// Store is an in-memory message store. All access is guarded by mu.
type Store struct {
	mu       sync.RWMutex
	messages map[int]Message
	maxID    int
}

func New() *Store {
	return &Store{messages: make(map[int]Message)}
}

// Add stores text and returns the new message. ID is only consumed here,
// i.e. after the caller has already decoded and validated successfully.
func (s *Store) Add(text string) Message {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.maxID++
	m := Message{ID: s.maxID, Text: text}
	s.messages[m.ID] = m
	return m
}

func (s *Store) Get(id int) (Message, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	m, ok := s.messages[id]
	return m, ok
}
