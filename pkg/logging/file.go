package logging

import (
	"fmt"
	"time"
)

// getLogFilePath get the log file save path
func getLogFilePath() string {
	return "./runtime/logs/"
}

// getLogFileName get the save name of the log file
func getLogFileName() string {
	return fmt.Sprintf("log%s.log",
		time.Now().Format("20060102"),
	)
}
