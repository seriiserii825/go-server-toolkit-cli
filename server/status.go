package server

func GetFreeStorage(total, used float64) float64 {
	return total - used
}

func GetDiskUsagePercentage(total, used float64) float64 {
	if total == 0 {
		return 0
	}
	return (used / total) * 100
}
