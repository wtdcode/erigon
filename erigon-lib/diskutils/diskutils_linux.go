//go:build linux

package diskutils

import (
	"fmt"
	"os"
	"syscall"

	"github.com/ledgerwatch/log/v3"
)

func MountPointForDirPath(dirPath string) string {
	// Get the file information for the directory path
	fileInfo, err := os.Stat(dirPath)
	if err != nil {
		log.Debug("[diskutils] Error getting file info for dir path:", dirPath, "Error:", err)
		return "/"
	}

	// Get the underlying system-specific information
	stat := fileInfo.Sys().(*syscall.Stat_t)

	// Get the device ID
	dev := stat.Dev

	// Get the symbolic link for the device ID in the /dev/disk/by-uuid directory
	linkPath := fmt.Sprintf("/dev/disk/by-uuid/%x", dev)
	link, err := os.Readlink(linkPath)
	if err != nil {
		log.Debug("[diskutils] Error getting symbolic link for device ID:", dev, "Error:", err)
		return "/"
	}

	// Print the mount point
	fmt.Println("Mount point:", link)
	log.Info("[diskutils] Mount point:", link)

	return link
}
