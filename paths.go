package state

import (
	"fmt"
	"os"
	"path/filepath"
)

func Path(dir string) string {
	base := ExeName()
	file := fmt.Sprintf("%s.sock", base)
	return filepath.Join(dir, file)
}

func PathWithPid(dir string) string {
	base := ExeName()
	file := fmt.Sprintf("%s.%d.sock", base, os.Getpid())
	return filepath.Join(dir, file)
}

func ExeName() string {
	exe, err := os.Executable()
	if err != nil {
		return os.Args[0]
	}
	return filepath.Base(exe)
}
