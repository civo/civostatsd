package gather

import (
	"time"

	"github.com/shirou/gopsutil/cpu"
)

func usedCPU() float64 {
	used, _ := cpu.Percent(time.Duration(5)*time.Second, false)
	return used[0]
}
