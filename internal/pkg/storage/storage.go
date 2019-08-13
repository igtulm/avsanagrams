package storage

import (
	"sync"
)

type storage struct {
	data   map[string][]string
	rwlock *sync.RWMutex
}

func New() Storage {
	s := &storage{
		data:   make(map[string][]string),
		rwlock: &sync.RWMutex{},
	}
	return s
}

func (st *storage) Add(key string, value string) {
	st.rwlock.Lock()
	defer st.rwlock.Unlock()

	st.data[key] = append(st.data[key], value)
}

func (st storage) Get(key string) []string {
	st.rwlock.RLock()
	defer st.rwlock.RUnlock()

	return st.data[key]
}
