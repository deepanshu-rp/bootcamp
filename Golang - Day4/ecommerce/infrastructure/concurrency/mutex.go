package concurrency

import (
	"sync"

	"github.com/google/uuid"
)

type mutex struct {
	mutex sync.Map
}

// Why create mutex here?
var Mutex mutex

// return true if can apply lock
func (m *mutex) Lock(id uuid.UUID) bool {
	_, found := m.mutex.Load(id)
	if found {
		return false
	}

	m.mutex.Store(id, true)
	return true
}

func (m *mutex) Unlock(id uuid.UUID) {
	m.mutex.Delete(id)
}
