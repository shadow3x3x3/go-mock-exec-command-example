package devices

import (
	"fmt"
	"go-mock-exec-command-example/util/cmd"
	"path/filepath"
	"strings"
)

// IsExists uses linux lsblk command for checking the device is exists or not.
func IsExists(targetDev string) bool {
	devPath := devPath(targetDev)

	if _, err := cmd.ExecutionCMD("lsblk " + devPath); err != nil {
		return false
	}

	return true
}

// GetSize get the size of targetDev by linux lsblk command
// The argument of lsblk command is "-lno SIZE"
func GetSize(targetDev string) (string, error) {
	devPath := devPath(targetDev)

	rawSize, err := cmd.ExecutionCMD("lsblk -lno SIZE " + devPath)
	if err != nil {
		return "", fmt.Errorf("Can not get the size of %s", targetDev)
	}

	return strings.TrimSpace(rawSize), nil
}

func devPath(dev string) string {
	return filepath.Join("/dev/", dev)
}
