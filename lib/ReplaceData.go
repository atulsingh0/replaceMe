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

	// getting Env var list
	for _, val := range os.Environ() {
		sList := strings.Split(val, "=")
		k := sList[0]
		v := sList[1]
		// fmt.Println(k, v)
		if !strings.HasPrefix(k, "_") {
			// TODO: Use os.Expand(s, os.Getenv) to replace all the env natively
			data = []byte(strings.Replace(string(data), k, v, -1))
		}
	}
	return data
}
