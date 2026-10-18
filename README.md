# nvidia-smi-json
nvidia-smi xml to json

```
go install github.com/fffaraz/nvidia-smi-json/cmd/nvidia-smi-json@latest

nvidia-smi-json -q
OR
nvidia-smi -q -x | nvidia-smi-json > output.json
```

## Library usage

```
go get github.com/fffaraz/nvidia-smi-json
```

```go
package main

import (
	"fmt"

	nvidiasmijson "github.com/fffaraz/nvidia-smi-json"
)

func main() {
	log := nvidiasmijson.GetNvidiaSmiLog() // nil if nvidia-smi is missing or fails
	if log == nil {
		fmt.Println("nvidia-smi not available")
		return
	}
	for _, gpu := range log.GPUS {
		fmt.Printf("%s: %s, util %s, mem %s / %s\n",
			gpu.ProductName,
			gpu.Temperature.GpuTemp,
			gpu.Utilization.GpuUtil,
			gpu.FbMemoryUsage.Used,
			gpu.FbMemoryUsage.Total)
	}
}
```
