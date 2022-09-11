package lib

import (
	"os"
	"strings"
)

func ReplaceDataFromEnv(data []byte) []byte {

	var out_data []byte
	// getting Env var list
	for _, val := range os.Environ() {
		sList := strings.Split(val, "=")
		k := sList[0]
		v := sList[1]
		// fmt.Println(k, v)
		if !strings.HasPrefix(k, "_") {
			out_data = Replace(data, k, v)
		}
		data = out_data
	}
	return data
}
