package gather

import "github.com/cloudfoundry/gosigar"

const output_format = "%-15s %4s %4s %5s %4s %-15s\n"

func formatSize(size uint64) string {
	return sigar.FormatSize(size * 1024)
}

func usedDisk() float32 {
	fslist := sigar.FileSystemList{}
	fslist.Get()

	for _, fs := range fslist.List {
		if fs.DirName == "/" {
			usage := sigar.FileSystemUsage{}
			usage.Get(fs.DirName)
			used := float32(usage.Used) / float32(usage.Total) * 100.0
			if used > 100.0 {
				return 100.0
			}
			return used
		}
	}
	return -1
}
