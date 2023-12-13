package services

import "fmt"

func MeetsThreshold(alertLevel int, time float64) bool {
	if alertLevel < 0 {
		panic("alertLevel must be positive")
	}

	switch alertLevel {
	case 0:
		return time >= 30
	case 1:
		return time >= 60
	case 2:
		return time >= 600
	default:
		return time >= float64(3600*(alertLevel-2))
	}
}

func GetAlertText(alertLevel int) string {
	if alertLevel < 0 {
		panic("alertLevel must be positive")
	}

	switch alertLevel {
	case 0:
		return "30秒"
	case 1:
		return "1分"
	case 2:
		return "10分"
	default:
		return fmt.Sprintf("%d時間", alertLevel-2)
	}
}
