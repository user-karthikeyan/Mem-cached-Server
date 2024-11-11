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
	)

	name, args, _ := strings.Cut(command, " ")

	switch name {
	case "set":
		parsed, _ = fmt.Sscanf(args, "%s %d %d %d $", &block.Key, &block.Flags, &block.Expiry, &block.Byte_count)
		block.Data_block, err = readFromClient(block.Byte_count)

		if err != nil {
			result = fmt.Sprintf("Error: %v\n", err)
			break
		} else if parsed == 4 {
			result = putBlock(block)
		} else {
			result = "Invalid arguments\r\n"
		}
	case "get":
		parsed, _ = fmt.Sscanf(args, "%s $", &key)

		if parsed == 1 {
			result = getBlock(key)
		} else {
			result = "Invalid arguments\r\n"
		}

	case "append":
		parsed, _ = fmt.Sscanf(args, "%s %d %d %d $", &block.Key, &block.Flags, &block.Expiry, &block.Byte_count)
		block.Data_block, err = readFromClient(block.Byte_count)

		if err != nil {
			result = fmt.Sprintf("Error: %v\n", err)
			break
		} else if parsed == 4 {
			result = appendBlock(block.Key, block)
		} else {
			result = "Invalid arguments\r\n"
		}

	case "prepend":
		parsed, _ = fmt.Sscanf(args, "%s %d %d %d $", &block.Key, &block.Flags, &block.Expiry, &block.Byte_count)
		block.Data_block, err = readFromClient(block.Byte_count)

		if err != nil {
			result = fmt.Sprintf("Error: %v\n", err)
			break
		} else if parsed == 4 {
			result = prependBlock(block.Key, block)
		} else {
			result = "Invalid arguments\r\n"
		}

	default:
		result = "Invalid Command!\r\n"
	}

	return result
}
