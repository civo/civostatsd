package gather

import (
	"time"

	"github.com/cloudfoundry/gosigar"
)

func usedCPU() float32 {
	cpu1 := sigar.Cpu{}
	cpu1.Get()
	time.Sleep(time.Second)
	cpu2 := sigar.Cpu{}
	cpu2.Get()
	cpu := cpu2.Delta(cpu1)
	busyTime := cpu.User + cpu.Nice + cpu.Sys
	totalTime := busyTime + cpu.Idle + cpu.Wait
	used := float32(busyTime) / float32(totalTime) * 100.0
	if used > 100.0 {
		return 100.0
	}
	return used
}
