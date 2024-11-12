package commands

import (
	"fmt"
	"strings"
)

func ParseCommand(command string, readFromClient func(int) (string, error)) string {

	var (
		block  Datablock
		parsed int
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
			result = putBlock(block)
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

	case "append":
		parsed, _ = fmt.Sscanf(args, "%s %d %d %d %s", &block.Key, &block.Flags, &block.Expiry, &block.Byte_count, &end)
		block.Data_block, err = readFromClient(block.Byte_count)

		if err != nil {
			result = fmt.Sprintf("Error: %v\n", err)
			break
		} else if parsed == 5 {
			result = appendBlock(block.Key, block)
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
			result = prependBlock(block.Key, block)
		} else {
			result = "Invalid arguments\r\n"
		}

	default:
		result = "Invalid Command!\r\n"
	}

	return result
}
