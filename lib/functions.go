package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func outputFilechk(inp string, out string) (string, string) {
	if out == "" {
		out = inp + ".bak"

		bytesRead, err := ioutil.ReadFile(inp)
		if err != nil {
			log.Fatal(err)
		}

		err = ioutil.WriteFile(out, bytesRead, 0644)
		if err != nil {
			log.Fatal(err)
		}
		inp, out = out, inp
		fmt.Printf("Processed file: %v\n", out)
		fmt.Printf("Backup input file is: %v\n", inp)
	} else {
		fmt.Printf("Input file: %v\n", inp)
		fmt.Printf("Processed file is: %v\n", out)
	}
	return inp, out
}

func replace(data []byte, key string, val string) []byte {
	return []byte(strings.Replace(string(data), key, val, -1))
}

func replaceData(data []byte, strMapToReplace string) []byte {

	var keyMap map[string]interface{}
	// Converting string to Map data type
	json.Unmarshal([]byte(strMapToReplace), &keyMap)

	for key, val := range keyMap {
		v := fmt.Sprintf("%v", val)
		out_data := replace(data, key, v)
		data = out_data
	}
	return data
}

func replaceDataFromEnv(data []byte) []byte {

	var out_data []byte
	// getting Env var list
	for _, val := range os.Environ() {
		sList := strings.Split(val, "=")
		k := sList[0]
		v := sList[1]
		// fmt.Println(k, v)
		if !strings.HasPrefix(k, "_") {
			out_data = replace(data, k, v)
		}
		data = out_data
	}
	return data
}
