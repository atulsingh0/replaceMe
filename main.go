package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	var (
		inputFile          string
		stringToSearch     string
		stringToReplace    string
		stringToReplaceEnv string
		outputFile         string
		valueOut           string
	)

	flag.StringVar(&inputFile, `i`, "", "-i <Input_file_path>")
	flag.StringVar(&stringToSearch, `k`, "", "-k <String_to_search>")
	flag.StringVar(&stringToReplace, `v`, "", "-v <String_to_replace_with>")
	flag.StringVar(&stringToReplaceEnv, `e`, "", "-e <Environment_var_name_which_has_string_to_replace_with>")
	flag.StringVar(&outputFile, `o`, "", "Output file path")

	flag.Parse()

	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal("Input file: ", err)
		return
	}

	if stringToReplaceEnv != "" {
		valueOut = os.Getenv(stringToReplaceEnv)
	} else {
		valueOut = stringToReplace
	}

	out_data := strings.Replace(string(data), stringToSearch, valueOut, -1)
	ioutil.WriteFile(outputFile, []byte(out_data), 0777)
	if err != nil {
		log.Fatal("Output file: ", err)
		return
	}
}
