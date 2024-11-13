package commands

import (
	"fmt"
	"strings"
)

func ParseCommand(command string, readFromClient func(int) (string, error)) string {

	var (
		block  Datablock
		parsed int
		x      int
		result string
		key    string
		err    error
		end    string
	)

	name, args, _ := strings.Cut(command, " ")

	switch name {
	case "set":
		parsed, _ = fmt.Sscanf(args, "%s %d %d %d %s", &block.Key, &block.Flags, &block.Expiry, &block.Byte_count, &end)
		block.Data_block, err = readFromClient(block.Byte_count)

		if err != nil {
			result = fmt.Sprintf("Error: %v\n", err)
			break
		} else if parsed == 5 {
			result = putBlock(&block)
		} else {
			result = "Invalid arguments\r\n"
		}
	case "cas":
		parsed, _ = fmt.Sscanf(args, "%s %d %d %d %d %s", &block.Key, &block.Flags, &block.Expiry, &block.Byte_count, &block.CAS, &end)
		block.Data_block, err = readFromClient(block.Byte_count)

		if err != nil {
			result = fmt.Sprintf("Error: %v\n", err)
			break
		} else if parsed == 6 {
			result = checkAndSetBlock(&block)
		} else {
			result = "Invalid arguments\r\n"
		}
	case "get":
		parsed, _ = fmt.Sscanf(args, "%s %s", &key, &end)

		if parsed == 2 {
			result = getBlock(key)
		} else {
			result = "Invalid arguments\r\n"
		}
	case "gets":
		parsed, _ = fmt.Sscanf(args, "%s %s", &key, &end)

		if parsed == 2 {
			result = getsBlock(key)
		} else {
			result = "Invalid arguments\r\n"
		}
	case "delete":
		parsed, _ = fmt.Sscanf(args, "%s %s", &key, &end)

		if parsed == 2 {
			result = deleteBlock(key)
		} else {
			result = "Invalid arguments\r\n"
		}
	case "incr":
		parsed, _ = fmt.Sscanf(args, "%s %d %s", &key, &x, &end)

		if parsed == 3 {
			result = increment(key, x)
		} else {
			result = "Invalid arguments\r\n"
		}
	case "decr":
		parsed, _ = fmt.Sscanf(args, "%s %d %s", &key, &x, &end)

		if parsed == 3 {
			result = decrement(key, x)
		} else {
			result = "Invalid arguments\r\n"
		}
	case "append":
		parsed, _ = fmt.Sscanf(args, "%s %d %d %d %s", &block.Key, &block.Flags, &block.Expiry, &block.Byte_count, &end)
		block.Data_block, err = readFromClient(block.Byte_count)

		if err != nil {
			result = fmt.Sprintf("Error: %v\n", err)
			break
		} else if parsed == 5 {
			result = appendBlock(block.Key, &block)
		} else {
			result = "Invalid arguments\r\n"
		}

	case "prepend":
		parsed, _ = fmt.Sscanf(args, "%s %d %d %d %s", &block.Key, &block.Flags, &block.Expiry, &block.Byte_count, &end)
		block.Data_block, err = readFromClient(block.Byte_count)

		if err != nil {
			result = fmt.Sprintf("Error: %v\n", err)
			break
		} else if parsed == 5 {
			result = prependBlock(block.Key, &block)
		} else {
			result = "Invalid arguments\r\n"
		}
	case "replace":
		parsed, _ = fmt.Sscanf(args, "%s %d %d %d %s", &block.Key, &block.Flags, &block.Expiry, &block.Byte_count, &end)
		block.Data_block, err = readFromClient(block.Byte_count)

		if err != nil {
			result = fmt.Sprintf("Error: %v\n", err)
			break
		} else if parsed == 5 {
			result = replaceBlock(block.Key, &block)
		} else {
			result = "Invalid arguments\r\n"
		}
	case "add":
		parsed, _ = fmt.Sscanf(args, "%s %d %d %d %s", &block.Key, &block.Flags, &block.Expiry, &block.Byte_count, &end)
		block.Data_block, err = readFromClient(block.Byte_count)

		if err != nil {
			result = fmt.Sprintf("Error: %v\n", err)
			break
		} else if parsed == 5 {
			result = addBlock(block.Key, &block)
		} else {
			result = "Invalid arguments\r\n"
		}
	default:
		result = "Invalid Command!\r\n"
	}

	return result
}
