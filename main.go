package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	repme "github.com/atulsingh0/replaceme/lib"
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

	// Reading file
	data, err := ioutil.ReadFile(inputFile)

	if err != nil {
		log.Fatal("Input file: ", err)
		return
	} else if outputFile == "" {
		outputFile = inputFile + ".bak"
		ioutil.WriteFile(outputFile, data, 0644)
		outputFile, inputFile = inputFile, outputFile
	}

	// Replacing the DATA based on strMapToReplace MAP
	out_data := repme.ReplaceData(data, strMapToReplace)
	ioutil.WriteFile(outputFile, out_data, 0600)

}
