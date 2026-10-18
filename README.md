# nvidia-smi-json

Convert `nvidia-smi` XML output to JSON.

## CLI

Install:

```sh
go install github.com/fffaraz/nvidia-smi-json/cmd/nvidia-smi-json@latest
```

Run `nvidia-smi` and print JSON directly:

```sh
nvidia-smi-json -q
```

Or pipe existing XML output:

```sh
nvidia-smi -q -x | nvidia-smi-json > output.json
```

## Library

Install:

```sh
go get github.com/fffaraz/nvidia-smi-json
```

Example:

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
