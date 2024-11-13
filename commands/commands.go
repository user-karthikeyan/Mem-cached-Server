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

func replaceBlock(key string, block *Datablock) string {
	lock := data_block.LoadLock(key)

	if lock != nil {
		defer lock.Unlock()
		lock.Lock()
		cache[key].Replace(block.Data_block, block.Byte_count, counter.GetValue())
		return "STORED\r\n"
	} else {
		return "NOT_STORED\r\n"
	}
}

func appendBlock(key string, block *Datablock) string {
	lock := data_block.LoadLock(key)

	if lock != nil {
		defer lock.Unlock()
		lock.Lock()
		cache[key].Append(block.Data_block, block.Byte_count, counter.GetValue())
		return "STORED\r\n"
	} else {
		return "NOT_STORED\r\n"
	}
}

func prependBlock(key string, block *Datablock) string {
	lock := data_block.LoadLock(key)

	if lock != nil {
		defer lock.Unlock()
		lock.Lock()
		cache[key].Prepend(block.Data_block, block.Byte_count, counter.GetValue())
		return "STORED\r\n"
	} else {
		return "NOT_STORED\r\n"
	}
}

func getBlock(key string) string {

	lock := data_block.LoadLock(key)

	if lock != nil {
		lock.RLock()
		value := cache[key]

		if !removeExpired(value) {
			lock.RUnlock()
			return fmt.Sprintf("VALUE %s %d %d\r\n%s\r\nEND\r\n", value.Key, value.Flags, value.Byte_count, value.Data_block)
		} else {
			lock.RUnlock()
			deleteBlock(key)
		}

	}
	return "END\r\n"
}

func getsBlock(key string) string {
	lock := data_block.LoadLock(key)

	if lock != nil {
		lock.RLock()
		value := cache[key]

		if !removeExpired(value) {
			lock.RUnlock()
			return fmt.Sprintf("VALUE %s %d %d %d\r\n%s\r\nEND\r\n", value.Key, value.Flags, value.Byte_count, value.CAS, value.Data_block)
		} else {
			lock.RUnlock()
			deleteBlock(key)
		}

	}
	return "END\r\n"
}

func checkAndSetBlock(block *Datablock) string {
	lock := data_block.LoadLock(block.Key)

	if lock != nil {
		defer lock.Unlock()
		lock.Lock()
		value := cache[block.Key]

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

	lock := data_block.GetLock(block.Key)

	defer lock.Unlock()
	lock.Lock()

	block.AddExpiry()
	block.CAS = counter.GetValue()
	cache[block.Key] = block

	return "STORED\r\n"
}

func addBlock(key string, block *Datablock) string {

	lock := data_block.LoadLock(key)

	if lock != nil {
		return "NOT_STORED\r\n"
	} else {
		return putBlock(block)
	}
}

func increment(key string, x int) string {
	lock := data_block.LoadLock(key)

	if lock != nil {
		defer lock.Unlock()
		lock.Lock()

		value := cache[key]
		n, err := strconv.Atoi(value.Data_block)

		if err != nil {
			return "INCREMENT ON A NON_NUMERIC VALUE!\r\n"
		} else {
			value.Data_block = strconv.Itoa(n + x)
			value.CAS = counter.GetValue()
			return value.Data_block + "\r\n"
		}

	} else {
		return "NOT_FOUND"
	}
}

func decrement(key string, x int) string {
	lock := data_block.LoadLock(key)

	if lock != nil {
		defer lock.Unlock()
		lock.Lock()

		value := cache[key]
		n, err := strconv.Atoi(value.Data_block)

		if err != nil {
			return "DECREMENT ON A NON_NUMERIC VALUE!\r\n"
		} else {
			value.Data_block = strconv.Itoa(n - x)
			value.CAS = counter.GetValue()
			return value.Data_block + "\r\n"
		}

	} else {
		return "NOT_FOUND"
	}
}

func deleteBlock(key string) string {
	lock := data_block.LoadLock(key)

	if lock != nil {
		defer lock.Unlock()
		lock.Lock()
		delete(cache, key)
		data_block.DeleteLock(key)
		return "DELETED\r\n"
	} else {
		return "NOT_FOUND\r\n"
	}
}

func removeExpired(block *Datablock) bool {

	if block.Expiry < time.Now().Unix() && block.Expiry != 0 {
		return true
	}
	return false
}
