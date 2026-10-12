package main

import (
	"github.com/seriiserii825/go-server-toolkit-cli/server"
	"github.com/seriiserii825/go-server-toolkit-cli/utils"
)

func main() {
	serverName := server.GetServerName()
	totalStorage := server.GetTotalStorage()
	usedStorage := server.GetUsedStorage()
	freeStorage := server.GetFreeStorage(totalStorage, usedStorage)

	diskUsage := server.GetDiskUsagePercentage(totalStorage, usedStorage)
	status := utils.GetStorageStatus(diskUsage)
	isCritical := utils.IsStorageCritical(diskUsage)

	utils.PrintStorageInfo(serverName,
		utils.FormatStorage(totalStorage),
		utils.FormatStorage(usedStorage),
		utils.FormatStorage(freeStorage),
		utils.FormatPercentage(diskUsage),
		status,
		isCritical,
	)
}
