package nvidiasmijson

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

func XmlToObject(input []byte) *NvidiaSmiLog {
	var data NvidiaSmiLog
	if err := xml.Unmarshal(input, &data); nil != err {
		fmt.Fprintln(os.Stderr, "Error unmarshalling from XML:", err)
		return nil
	}
	return &data
}

func XmlToJson(input []byte) string {
	data := XmlToObject(input)
	output, err := json.MarshalIndent(data, "", "  ")
	if nil != err {
		fmt.Fprintln(os.Stderr, "Error marshalling to JSON:", err)
		return ""
	}
	return string(output)
}

func HasNvidiaSmi() bool {
	_, err := exec.LookPath("nvidia-smi")
	return err == nil
}

func GetNvidiaSmiLog() *NvidiaSmiLog {
	if !HasNvidiaSmi() {
		return nil
	}
	return XmlToObject(RunNvidiaSmi())
}

func RunNvidiaSmi() []byte {
	cmd := "nvidia-smi"
	if runtime.GOOS == "windows" {
		cmd += ".exe"
	}
	out, err := exec.Command(cmd, "-q", "-x").Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error running nvidia-smi:", err)
		return nil
	}
	return out
}
