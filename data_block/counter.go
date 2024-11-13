package data_block

import (
	"sync"
)

type Counter struct {
	count int64
	lock  sync.Mutex
}

func (counter *Counter) GetValue() int64 {
	defer counter.lock.Unlock()
	counter.lock.Lock()
	counter.count++
	return counter.count
}
