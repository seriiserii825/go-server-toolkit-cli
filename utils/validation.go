package utils

func IsStorageCritical(usage float64, treshold float64) bool {
	return usage >= treshold
}

func GetStorageStatus(usage float64, treshold float64) string {
	if usage < treshold {
		return "OK"
	}
	if usage >= treshold && usage < 90 {
		return "WARNING"
	}
	return "CRITICAL"
}
