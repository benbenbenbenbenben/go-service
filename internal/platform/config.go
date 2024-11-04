package platform

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
)

// Update these constants with your own service configuration details.
// Replace the service name, display name, and description as needed.
// LogFileName is optional and is used to demonstrate an example service core logic
// that appends text to a file periodically.
const (
	ServiceName    = "go-service"
	ServiceDisplay = "Go Service"
	ServiceDesc    = "A service that appends 'Hello World' to a file every 5 minutes."
	LogFileName    = "go-service-log.txt"
)

func GetInstallDir() string {
	switch runtime.GOOS {
	case "darwin":
		return "/usr/local/opt/" + ServiceName
	case "linux":
		return "/opt/" + ServiceName
	case "windows":
		return filepath.Join(os.Getenv("ProgramData"), ServiceName)
	default:
		return ""
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
