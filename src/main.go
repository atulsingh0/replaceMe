package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func flagUsage() {
	fmt.Println("Usage:")
	fmt.Println("---------------\n")
	fmt.Println("./replaceMe -i <inp-file> -m <key-value-map>")
	fmt.Println("i.e - ./replaceMe -i input.txt -m '{\"name\" : \"johndoe\", \"age\": 20}'\n")
	fmt.Println("./replaceMe -i <inp-file> -o <out-file> -m <key-value-map>")
	fmt.Println("i.e - ./replaceMe -i input.txt -o output.txt -m '{\"name\" : \"johndoe\", \"age\": 20}'\n")
}

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

func main() {

	var (
		inputFile       string
		strMapToReplace string
		outputFile      string
	)

	flag.StringVar(&inputFile, `i`, "", "-i <Input_file_path>")
	flag.StringVar(&strMapToReplace, `m`, "", "-m <key-value-map>")
	flag.StringVar(&outputFile, `o`, "", "Output file path")

	flag.Parse()

	if inputFile == "" {
		flag.Usage()
		log.Fatal("-i flag is required.\n")
		os.Exit(1)
	}

	inp, out := outputFilechk(inputFile, outputFile)

	// Reading file
	data, err := ioutil.ReadFile(inp)
	if err != nil {
		log.Fatal("Input file: ", err)
		return
	} else {
		// Replacing the DATA based on strMapToReplace MAP
		out_data := replaceData(data, strMapToReplace)
		out_data = replaceDataFromEnv(out_data)
		ioutil.WriteFile(out, out_data, 0600)
	}
}
