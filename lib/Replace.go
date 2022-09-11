package lib

import "strings"

func Replace(data []byte, key string, val string) []byte {
	return []byte(strings.Replace(string(data), key, val, -1))
}
