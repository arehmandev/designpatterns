package singleton

import "sync"

// Singleton - done
type Singleton map[string]string

var (
	once sync.Once

	instance Singleton
)

// New - instantiates a singleton
func New() Singleton {
	once.Do(func() {
		instance = make(Singleton)
	})

	return instance
}
