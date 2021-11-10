package state

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

func SingletonPath() string {
	dir := SocketFolder()
	name := ExecutableName()
	file := fmt.Sprintf("%s.state", name)
	return filepath.Join(dir, file)
}

func InstancePath() string {
	pid := os.Getpid()
	dir := SocketFolder()
	name := ExecutableName()
	file := fmt.Sprintf("%s.%d.state", name, pid)
	return filepath.Join(dir, file)
}

func SocketFolder() string {
	if runtime.GOOS == "windows" {
		return ExecutableFolder()
	}
	switch runtime.GOOS {
	case "windows":
		return ExecutableFolder()
	case "linux", "darwin":
		name := ExecutableName()
		file := fmt.Sprintf("%s.state", name)
		path := filepath.Join("/var/run", file)
		var fid, err = os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0755)
		if err != nil {
			return ExecutableFolder()
		}
		defer fid.Close()
		return "/var/run"
	default:
		log.Fatalf("Unknown GOOS %s", runtime.GOOS)
		return ExecutableFolder()
	}
}

func ExecutablePath() string {
	path, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	return path
}

func ExecutableFolder() string {
	path := ExecutablePath()
	return filepath.Dir(path)
}

func ExecutableName() string {
	path := ExecutablePath()
	return filepath.Base(path)
}

func ExecutableWithExt(ext string) string {
	path := ExecutablePath()
	dir := filepath.Dir(path)
	base := filepath.Base(path)
	file := base + "." + ext
	return filepath.Join(dir, file)
}
