package replaceme

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func ReplaceData(data []byte, strMapToReplace string) []byte {

	var keyMap map[string]interface{}
	// Converting string to Map data type
	json.Unmarshal([]byte(strMapToReplace), &keyMap)

	// Creating Env var from Map
	for key, val := range keyMap {
		os.Setenv(key, fmt.Sprintf("%v", val))
	}

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

func Replace(data []byte, key string, val string) []byte {
	return []byte(strings.Replace(string(data), key, val, -1))
}
