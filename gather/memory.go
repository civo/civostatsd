package gather

import "github.com/cloudfoundry/gosigar"

func format(val uint64) uint64 {
	return val / 1024
}

func usedMemory() float32 {
	mem := sigar.Mem{}
	mem.Get()
	used := float32(mem.ActualUsed) / float32(mem.Total) * 100.0
	if used > 100.0 {
		return 100.0
	}
	return used
}
