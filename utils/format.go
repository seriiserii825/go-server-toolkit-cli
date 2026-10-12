package utils

import (
	"fmt"

	"github.com/fatih/color"
)

func FormatStorage(value float64) string {
	if value >= 1e12 {
		return fmt.Sprintf("%.2f TB", value/1e12)
	} else if value >= 1e9 {
		return fmt.Sprintf("%.2f GB", value/1e9)
	} else if value >= 1e6 {
		return fmt.Sprintf("%.2f MB", value/1e6)
	} else if value >= 1e3 {
		return fmt.Sprintf("%.2f KB", value/1e3)
	}
	return fmt.Sprintf("%.2f B", value)
}

func FormatPercentage(value float64) string {
	return fmt.Sprintf("%.2f%%", value)
}

func PrintStorageInfo(serverName, total, used, free, diskUsage, status string, isCritical bool) {
	intro := `SERVER TOOLKIT
================================
	`

	result := fmt.Sprintf(`%s

	Server Name: %s

	Total Storage: %s
	Used Storage: %s
	Free Storage: %s

	Disk Usage: %s
	Status: %s
	`, intro, serverName, total, used, free, diskUsage, status)

	fmt.Println(result)

	if isCritical {
		color.Red(result)
	}
}
