package data_block

import (
	"time"
)

type Datablock struct {
	Key        string
	Data_block string
	Flags      uint16
	Expiry     int64
	Byte_count int
}

func (block *Datablock) Append(Data_block string, Byte_count int) {
	block.Data_block += Data_block
	block.Byte_count += Byte_count
}

func (block *Datablock) Prepend(Data_block string, Byte_count int) {
	block.Data_block = Data_block + block.Data_block
	block.Byte_count += Byte_count
}

func (block *Datablock) AddExpiry() {
	if block.Expiry != 0 {
		block.Expiry += time.Now().Unix()
	}
}
