package kyoto

import (
	"net/http"
	"sync"
)

type State interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Del(key string)
	Export() map[string]interface{}
}

type Context interface {
	Get(key string) interface{}
	Set(key string, value interface{})
	Del(key string)
	GetResponseWriter() http.ResponseWriter
	GetRequest() *http.Request
}

type Store struct {
	state map[string]interface{}
	lock  sync.RWMutex
}

func NewStore() *Store {
	return &Store{state: make(map[string]interface{})}
}

func (s *Store) Get(key string) interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.state[key]
}

func (s *Store) Set(key string, value interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.state[key] = value
}

func (s *Store) Del(key string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	delete(s.state, key)
}

func (s *Store) Export() map[string]interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.state
}

// Aliases to avoid verbose type casting

func (s *Store) GetResponseWriter() http.ResponseWriter {
	return s.Get("internal:rw").(http.ResponseWriter)
}

func (s *Store) GetRequest() *http.Request {
	return s.Get("internal:r").(*http.Request)
}