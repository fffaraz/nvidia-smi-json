package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	nvidiasmijson "github.com/fffaraz/nvidia-smi-json"
)

func main() {
	var query bool
	flag.BoolVar(&query, "q", false, "run nvidia-smi -q")
	flag.Parse()
	if query {
		if nvidiasmijson.HasNvidiaSmi() {
			fmt.Println(nvidiasmijson.XmlToJson(nvidiasmijson.RunNvidiaSmi()))
		} else {
			fmt.Fprintln(os.Stderr, "nvidia-smi not found")
		}
		return
	}
	input, err := ioutil.ReadAll(os.Stdin)
	if nil != err {
		fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
		return
	}
	fmt.Println(nvidiasmijson.XmlToJson(input))
}
