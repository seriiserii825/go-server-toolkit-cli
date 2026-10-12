package utils

import (
	"fmt"

	"github.com/fatih/color"
)

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
		color.Red("WARNING: Disk usage is critical!")
	}
}
