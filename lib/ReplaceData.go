package lib

import (
	"encoding/json"
	"fmt"
)

func ReplaceData(data []byte, strMapToReplace string) []byte {

	var keyMap map[string]interface{}
	// Converting string to Map data type
	json.Unmarshal([]byte(strMapToReplace), &keyMap)

	for key, val := range keyMap {
		v := fmt.Sprintf("%v", val)
		out_data := Replace(data, key, v)
		data = out_data
	}
	return data
}
