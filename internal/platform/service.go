package platform

import (
	"fmt"
	"runtime"
)

type Service interface {
	Install(execPath string) error
	Uninstall() error
	Status() (bool, error)
	Start() error
	Stop() error
}

func NewService() (Service, error) {
	switch runtime.GOOS {
	case "windows":
		return &windowsService{}, nil
	case "linux":
		return &linuxService{}, nil
	case "darwin":
		return &darwinService{}, nil
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}
