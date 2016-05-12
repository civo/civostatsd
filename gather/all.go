package gather

import "fmt"

type Stats struct {
	CPU    float32
	Disk   float32
	Memory float32
}

func All() Stats {
	stats := Stats{
		CPU:    usedCPU(),
		Disk:   usedDisk(),
		Memory: usedMemory(),
	}
	return stats
}

func (s Stats) String() string {
	return fmt.Sprintf("CPU: %.2f%%, Disk: %.2f%%, Memory: %.2f%%", s.CPU, s.Disk, s.Memory)
}
