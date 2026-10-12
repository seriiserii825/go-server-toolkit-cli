package utils

const (
	WarningThreshold  = 70.0
	CriticalThreshold = 90.0
)

func IsStorageCritical(usage float64) bool {
	return usage >= CriticalThreshold
}

func GetStorageStatus(usage float64) string {
	if usage < WarningThreshold {
		return "OK"
	}
	if usage < CriticalThreshold {
		return "WARNING"
	}
	return "CRITICAL"
}
