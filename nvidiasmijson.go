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
	if data.GPUS == nil {
		data.GPUS = []GPU{}
	}
	for i := range data.GPUS {
		gpu := &data.GPUS[i]
		// Remove verbose supported clocks list
		gpu.SupportedClocks = []SupportedMemClock{}
		// Encode empty lists as [] instead of null
		if gpu.MigDevices == nil {
			gpu.MigDevices = []MigDevice{}
		}
		if gpu.Processes == nil {
			gpu.Processes = []ProcessInfo{}
		}
		if gpu.AccountedProcesses == nil {
			gpu.AccountedProcesses = []AccountedProcess{}
		}
	}
	return &data
}

func XmlToJson(input []byte) string {
	data := XmlToObject(input)
	if data == nil {
		return ""
	}
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
