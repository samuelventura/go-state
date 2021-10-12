package state

import (
	"fmt"
	"os"
	"path/filepath"
)

func SingletonPath(dir string) string {
	name := ExecutableName()
	file := fmt.Sprintf("%s.state", name)
	return filepath.Join(dir, file)
}

func InstancePath(dir string, id int) string {
	name := ExecutableName()
	file := fmt.Sprintf("%s.%d.state", name, id)
	return filepath.Join(dir, file)
}

func ExecutableName() string {
	exe, err := os.Executable()
	if err != nil {
		return os.Args[0]
	}
	return filepath.Base(exe)
}
