package main

import (
	"fmt"
	"log/slog"
	"sync"
)

type Storage struct {
	storage map[string]string
	mu      sync.Mutex
}

func NewStorage() Storage {
	return Storage{
		storage: make(map[string]string),
	}
}

func (s *Storage) StoreIfNotExists(key, value string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.storage[key]; !exists {
		slog.Info("store element")
		s.storage[key] = value
	} else {
		slog.Info("element already exists")
	}
}

func main() {
	storage := NewStorage()
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			storage.StoreIfNotExists("my-key", fmt.Sprintf("my-value-%d", i))
		}()
	}

	wg.Wait()
}
