package gather

import (
	"github.com/shirou/gopsutil/disk"
)

func usedDisk() float32 {
	fslist, _ := disk.Partitions(false)

	for _, fs := range fslist {
		if fs.Mountpoint == "/" {
			usage, _ := disk.Usage(fs.Mountpoint)
			used := usage.UsedPercent
			return float32(used)
		}
	}
	return -1
}
