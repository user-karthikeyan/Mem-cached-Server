package commands

import (
	"MEM-CACHED-SERVER/data_block"
	"fmt"
	"strconv"
	"time"
)

type Datablock = data_block.Datablock
type Counter = data_block.Counter

var cache map[string]*Datablock = make(map[string]*Datablock)
var counter Counter

func appendBlock(key string, block *Datablock) string {
	_, present := cache[key]

	if present {
		cache[key].Append(block.Data_block, block.Byte_count, counter.GetValue())
		return "STORED\r\n"
	} else {
		return "NOT_STORED\r\n"
	}
}

func prependBlock(key string, block *Datablock) string {
	_, present := cache[key]

	if present {
		cache[key].Prepend(block.Data_block, block.Byte_count, counter.GetValue())
		return "STORED\r\n"
	} else {
		return "NOT_STORED\r\n"
	}
}

func getBlock(key string) string {

	value, present := cache[key]

	if present && !removeExpired(value) {
		defer value.Lock.RUnlock()
		value.Lock.RLock()
		return fmt.Sprintf("VALUE %s %d %d\r\n%s\r\nEND\r\n", value.Key, value.Flags, value.Byte_count, value.Data_block)
	} else {
		return "END\r\n"
	}
}

func getsBlock(key string) string {

	value, present := cache[key]

	if present && !removeExpired(value) {
		defer value.Lock.RUnlock()
		value.Lock.RLock()
		return fmt.Sprintf("VALUE %s %d %d %d\r\n%s\r\nEND\r\n", value.Key, value.Flags, value.Byte_count, value.CAS, value.Data_block)
	} else {
		return "END\r\n"
	}
}

func checkAndSetBlock(block *Datablock) string {
	value, present := cache[block.Key]

	if present {
		defer value.Lock.Unlock()
		value.Lock.Lock()

		if block.CAS == value.CAS {
			cache[block.Key] = block
			return "STORED\r\n"
		} else {
			return "EXISTS\r\n"
		}
	} else {
		return "NOT_FOUND\r\n"
	}
}

func putBlock(block *Datablock) string {
	value, present := cache[block.Key]

	block.CAS = counter.GetValue()

	if present {
		defer value.Lock.Unlock()
		value.Lock.Lock()
		cache[block.Key] = block
	} else {
		block.AddExpiry()
		cache[block.Key] = block
	}

	return "STORED\r\n"
}

func addBlock(key string, block *Datablock) string {
	_, present := cache[key]

	if present {
		return "NOT_STORED\r\n"
	} else {
		block.CAS = counter.GetValue()
		cache[key] = block
		return "STORED\r\n"
	}
}

func replaceBlock(key string, block *Datablock) string {
	value, present := cache[key]

	if present {
		value.Replace(block.Data_block, block.Byte_count, counter.GetValue())
		return "STORED\r\n"
	} else {
		return "NOT_STORED\r\n"
	}
}

func increment(key string, x int) string {
	value, present := cache[key]

	if present {
		n, err := strconv.Atoi(value.Data_block)

		if err != nil {
			return "INCREMENT ON A NON_NUMERIC VALUE!\r\n"
		} else {
			value.Data_block = strconv.Itoa(n + x)
			return value.Data_block + "\r\n"
		}
	} else {
		return "NOT_FOUND\r\n"
	}
}

func decrement(key string, x int) string {
	value, present := cache[key]

	if present {
		n, err := strconv.Atoi(value.Data_block)

		if err != nil {
			return "DECREMENT ON A NON_NUMERIC VALUE!\r\n"
		} else {
			value.Data_block = strconv.Itoa(n - x)
			return value.Data_block + "\r\n"
		}
	} else {
		return "NOT_FOUND"
	}
}

func deleteBlock(key string) string {
	value, present := cache[key]

	if present {
		defer value.Lock.Unlock()
		value.Lock.Lock()
		delete(cache, key)
		return "DELETED\r\n"
	} else {
		return "NOT_FOUND\r\n"
	}
}

func removeExpired(block *Datablock) bool {

	if block.Expiry < time.Now().Unix() && block.Expiry != 0 {
		delete(cache, block.Key)
		return true
	}
	return false
}
