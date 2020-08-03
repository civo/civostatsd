package gather

import "fmt"

// Stats is to save all stats
type Stats struct {
	CPU    float64
	Disk   float32
	Memory float32
}

// All function to get all stats from the system
func All() Stats {

	stats := Stats{
		CPU:    usedCPU(),
		Disk:   usedDisk(),
		Memory: usedMemory(),
	}
	return stats
}

func (s Stats) String() string {
	return fmt.Sprintf("CPU: %.2f%%, Disk: %.2f%%, Memory: %.2f%% - v2.0", s.CPU, s.Disk, s.Memory)
}
