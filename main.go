package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func xml2json(input []byte) string {
	var data NvidiaSmiLog

	if err := xml.Unmarshal(input, &data); nil != err {
		fmt.Fprintln(os.Stderr, "Error unmarshalling from XML:", err)
		return ""
	}

	output, err := json.MarshalIndent(&data, "", "  ")
	if nil != err {
		fmt.Fprintln(os.Stderr, "Error marshalling to JSON:", err)
		return ""
	}

	return string(output)
}

func main() {
	var query bool
	flag.BoolVar(&query, "q", false, "run nvidia-smi -q")
	flag.Parse()
	if query {
		cmd := "nvidia-smi"
		if runtime.GOOS == "windows" {
			cmd += ".exe"
		}
		out, err := exec.Command(cmd, "-q", "-x").Output()
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error running nvidia-smi:", err)
			return
		}
		fmt.Println(xml2json(out))
	} else {
		input, err := ioutil.ReadAll(os.Stdin)
		if nil != err {
			fmt.Fprintln(os.Stderr, "Error reading from stdin:", err)
			return
		}
		fmt.Println(xml2json(input))
	}
}
