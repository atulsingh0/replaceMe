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
		out = inp
		inp = "out-" + inp
		os.Rename(out, inp)
	}
	return inp, out
}

func replaceData(data []byte, strMapToReplace string) []byte {

	var keyMap map[string]interface{}
	// Converting string to Map data type
	json.Unmarshal([]byte(strMapToReplace), &keyMap)

	for key, val := range keyMap {
		v := fmt.Sprintf("%v", val)
		out_data := []byte(strings.Replace(string(data), key, v, -1))
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

	if inputFile == "" && strMapToReplace == "" {
		flagUsage()
		log.Fatal("-i and -m flags are required.\n")
		os.Exit(1)
	} else if inputFile == "" {
		flag.Usage()
		log.Fatal("-i flag is required.\n")
		os.Exit(1)
	} else if strMapToReplace == "" {
		flag.Usage()
		log.Fatal("-m flag is required.\n")
		os.Exit(1)
	}

	inp, out := outputFilechk(inputFile, outputFile)
	fmt.Println(inp, out)
	// Reading file
	data, err := ioutil.ReadFile(inp)
	if err != nil {
		log.Fatal("Input file: ", err)
		return
	} else {
		out_data := replaceData(data, strMapToReplace)
		ioutil.WriteFile(out, out_data, 0600)
	}
}
