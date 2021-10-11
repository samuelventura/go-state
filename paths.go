package state

import (
	"fmt"
	"os"
	"path/filepath"
)

func Path(dir string) (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	base := filepath.Base(exe)
	file := fmt.Sprintf("%s.sock", base)
	return filepath.Join(dir, file), nil
}

func PathWithPid(dir string) (string, error) {
	exe, err := os.Executable()
	if err != nil {
		return "", err
	}
	base := filepath.Base(exe)
	file := fmt.Sprintf("%s.%d.sock", base, os.Getpid())
	return filepath.Join(dir, file), nil
}
