package store

import (
	"sync"
)

type Message struct {
	Id       string `json:"id"`
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Text     string `json:"text"`
}
type Store struct {
	mutex    sync.RWMutex
	storeMap MapType
}
type MapType map[string]*Message

func (s *Store) Get(id string) (*Message, bool) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	msg, ok := s.storeMap[id]
	return msg, ok
}

func (s *Store) Put(msg *Message) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	s.storeMap[msg.Id] = msg
}

func New() *Store {
	return &Store{
		storeMap: make(MapType),
	}
}
