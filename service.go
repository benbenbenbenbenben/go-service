package goservice

import (
	"fmt"
	"runtime"

	"github.com/benbenbenbenbenben/goservice/platform"
)

type Service interface {
	Install() error
	Uninstall() error
	Status() (bool, error)
	Start() error
	Stop() error
	GetInstallDir() string
	ServiceName() string
	ServiceDisplayName() string
	ServiceDescription() string
}

func NewService(executablePath string, serviceDisplayName string, serviceDesc string) (Service, error) {
	switch runtime.GOOS {
	case "windows":
		return &platform.WindowsService{
			ExecutablePath: executablePath,
			ServiceDisplay: serviceDisplayName,
			ServiceDesc:    serviceDesc,
		}, nil
	case "linux":
		return &platform.LinuxService{
			ExecutablePath: executablePath,
			ServiceDisplay: serviceDisplayName,
			ServiceDesc:    serviceDesc,
		}, nil
	case "darwin":
		return &platform.DarwinService{
			ExecutablePath: executablePath,
			ServiceDisplay: serviceDisplayName,
			ServiceDesc:    serviceDesc,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}
