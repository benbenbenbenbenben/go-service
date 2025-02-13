package platform

import (
	"fmt"
	"io"
	"os"
	"runtime"
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
		return &windowsService{
			executablePath: executablePath,
			serviceDisplay: serviceDisplayName,
			serviceDesc:    serviceDesc,
		}, nil
	case "linux":
		return &linuxService{
			executablePath: executablePath,
			serviceDisplay: serviceDisplayName,
			serviceDesc:    serviceDesc,
		}, nil
	case "darwin":
		return &darwinService{
			executablePath: executablePath,
			serviceDisplay: serviceDisplayName,
			serviceDesc:    serviceDesc,
		}, nil
	default:
		return nil, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}
}

func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer source.Close()

	destination, err := os.OpenFile(dst, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return err
}
