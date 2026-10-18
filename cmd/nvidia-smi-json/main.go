package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	nvidiasmijson "github.com/fffaraz/nvidia-smi-json"
)

func main() {
	var query bool
	flag.BoolVar(&query, "q", false, "run nvidia-smi -q -x and convert its output")
	flag.Parse()
	if query {
		if !nvidiasmijson.HasNvidiaSmi() {
			fmt.Fprintln(os.Stderr, "nvidia-smi not found")
			os.Exit(1)
		}
		input := nvidiasmijson.RunNvidiaSmi()
		if input == nil {
			os.Exit(1)
		}
		printJson(input)
		return
	}
	input, err := io.ReadAll(os.Stdin)
	if nil != err {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		os.Exit(1)
	}
	printJson(input)
}

func printJson(input []byte) {
	output := nvidiasmijson.XmlToJson(input)
	if output == "" {
		os.Exit(1)
	}
	fmt.Println(output)
}
