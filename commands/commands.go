package commands

import (
	"fmt"
	"time"
	"MEM-CACHED-SERVER/data_block"
)

type Datablock = data_block.Datablock
var cache map[string]*Datablock = make(map[string]*Datablock)

func appendBlock(key string, block Datablock) string {
	_, present := cache[key]

	if present {
		cache[key].Append(block.Data_block, block.Byte_count)
		return "STORED\r\n"
	} else {
		return "NOT_STORED\r\n"
	}
}

func prependBlock(key string, block Datablock) string {
	_, present := cache[key]

	if present {
		cache[key].Prepend(block.Data_block, block.Byte_count)
		return "STORED\r\n"
	} else {
		return "NOT_STORED\r\n"
	}
}

func getBlock(key string) string {

	value, present := cache[key]

	if present && !removeExpired(*value) {
		return fmt.Sprintf("VALUE %s %d %d\r\n%s\r\nEND\r\n", value.Key, value.Flags, value.Byte_count, value.Data_block)
	} else {
		return "END\r\n"
	}
}

func putBlock(block Datablock) string {
	block.AddExpiry()
	cache[block.Key] = &block
	return "STORED\r\n"
}

func removeExpired(block Datablock) bool {

	if block.Expiry < time.Now().Unix() && block.Expiry != 0 {
		delete(cache, block.Key)
		return true
	}
	return false
}
