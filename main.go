package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if nil != err {
		fmt.Println("Error reading from stdin:", err)
		return
	}

	var data NvidiaSmiLog

	if err := xml.Unmarshal(input, &data); nil != err {
		fmt.Println("Error unmarshalling from XML:", err)
		return
	}

	output, err := json.MarshalIndent(&data, "", "  ")
	if nil != err {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	fmt.Println(string(output))
}
