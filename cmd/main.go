package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/atulsingh0/reaplaceme/lib"
)

func flagUsage() {
	fmt.Println("Usage:")
	fmt.Println("---------------\n")
	fmt.Println("./replaceMe -i <inp-file> -m <key-value-map>")
	fmt.Println("i.e - ./replaceMe -i input.txt -m '{\"name\" : \"johndoe\", \"age\": 20}'\n")
	fmt.Println("./replaceMe -i <inp-file> -o <out-file> -m <key-value-map>")
	fmt.Println("i.e - ./replaceMe -i input.txt -o output.txt -m '{\"name\" : \"johndoe\", \"age\": 20}'\n")
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
