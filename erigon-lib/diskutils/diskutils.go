//go:build !darwin && !linux

package diskutils

import (
	"github.com/ledgerwatch/log/v3"
)
å
func MountPointForDirPath(dirPath string) string {
	log.Info("[diskutils] Implemented only for darwin")
	return "/"
}å
