package data_block

import (
	"sync"
	"time"
)

type Datablock struct {
	Key        string
	CAS        int64
	Data_block string
	Flags      uint16
	Expiry     int64
	Byte_count int
	Lock       sync.RWMutex
}

func (block *Datablock) Append(Data_block string, Byte_count int, CAS int64) {
	defer block.Lock.Unlock()
	block.Lock.Lock()
	block.CAS = CAS
	block.Data_block += Data_block
	block.Byte_count += Byte_count
}

func (block *Datablock) Prepend(Data_block string, Byte_count int, CAS int64) {
	defer block.Lock.Unlock()
	block.Lock.Lock()
	block.CAS = CAS
	block.Data_block = Data_block + block.Data_block
	block.Byte_count += Byte_count
}

func (block *Datablock) AddExpiry() {
	if block.Expiry != 0 {
		block.Expiry += time.Now().Unix()
	}
}

func (block *Datablock) Replace(Data_block string, Byte_count int, CAS int64) {
	defer block.Lock.Unlock()
	block.Lock.Lock()
	block.CAS = CAS
	block.Data_block = Data_block
	block.Byte_count = Byte_count
}
