package gather

import "github.com/shirou/gopsutil/mem"

func usedMemory() float32 {
	mem, _ := mem.VirtualMemory()
	used := mem.UsedPercent
	return float32(used)
}
