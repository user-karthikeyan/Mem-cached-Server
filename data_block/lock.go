package data_block

import "sync"

var locks map[string]*sync.RWMutex = make(map[string]*sync.RWMutex)
var globalLock sync.Mutex

func GetLock(key string) *sync.RWMutex {
	defer globalLock.Unlock()
	globalLock.Lock()

	lock, present := locks[key]

	if present {
		return lock
	} else {
		locks[key] = new(sync.RWMutex)
		return locks[key]
	}
}

func LoadLock(key string) *sync.RWMutex {
	defer globalLock.Unlock()
	globalLock.Lock()

	lock, present := locks[key]

	if present {
		return lock
	} else {
		return nil
	}
}

func DeleteLock(key string) {
	defer globalLock.Unlock()
	globalLock.Lock()
	delete(locks, key)
}
