package utils

import "fmt"

func FormatStorage(value float64) string {
	return fmt.Sprintf("%.2f GB", value)
}

func FormatPercentage(value float64) string {
	return fmt.Sprintf("%.2f%%", value)
}
