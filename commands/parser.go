package commands

import (
	"fmt"
	"strings"
)

func ParseCommand(command string, data_block string) string {

	var (
		block  Datablock
		parsed int
		result string
		key    string
	)

	block.Data_block = data_block
	name, args, _ := strings.Cut(command, " ")

	switch name {
	case "set":
		parsed, _ = fmt.Sscanf(args, "%s %d %d %d $", &block.Key, &block.Flags, &block.Expiry, &block.Byte_count)

		if parsed == 4 {
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

		if parsed == 4 {
			result = appendBlock(block.Key, block)
		} else {
			result = "Invalid arguments\r\n"
		}

	case "prepend":
		parsed, _ = fmt.Sscanf(args, "%s %d %d %d $", &block.Key, &block.Flags, &block.Expiry, &block.Byte_count)

		if parsed == 4 {
			result = prependBlock(block.Key, block)
		} else {
			result = "Invalid arguments\r\n"
		}

	default:
		result = "Invalid Command!\r\n"
	}

	return result
}
