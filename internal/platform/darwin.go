package platform

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type darwinService struct{}

const plistTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>%s</string>
    <key>ProgramArguments</key>
    <array>
        <string>%s</string>
        <string>-run</string>
    </array>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
    <key>WorkingDirectory</key>
    <string>%s</string>
</dict>
</plist>`

func (s *darwinService) Install(execPath string) error {
	installDir := GetInstallDir()
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return fmt.Errorf("failed to create installation directory: %w", err)
	}

	// Copy binary to installation directory
	installedBinary := filepath.Join(installDir, "bin", filepath.Base(execPath))
	if err := os.MkdirAll(filepath.Dir(installedBinary), 0755); err != nil {
		return fmt.Errorf("failed to create bin directory: %w", err)
	}

	if err := copyFile(execPath, installedBinary); err != nil {
		return fmt.Errorf("failed to copy binary: %w", err)
	}

	plistPath := filepath.Join("/Library/LaunchDaemons", ServiceName+".plist")
	content := fmt.Sprintf(plistTemplate, ServiceName, installedBinary, installDir)

	if err := os.WriteFile(plistPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write plist file: %w", err)
	}

	if err := exec.Command("launchctl", "load", plistPath).Run(); err != nil {
		return fmt.Errorf("failed to load service: %w", err)
	}
	return nil
}

func (s *darwinService) Uninstall() error {
	plistPath := filepath.Join("/Library/LaunchDaemons", ServiceName+".plist")

	if err := exec.Command("launchctl", "unload", plistPath).Run(); err != nil {
		return fmt.Errorf("failed to unload service: %w", err)
	}

	if err := os.Remove(plistPath); err != nil {
		return fmt.Errorf("failed to remove plist file: %w", err)
	}
	return nil
}

func (s *darwinService) Status() (bool, error) {
	err := exec.Command("launchctl", "list", ServiceName).Run()
	return err == nil, nil
}

func (s *darwinService) Start() error {
	if err := exec.Command("launchctl", "start", ServiceName).Run(); err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}
	return nil
}

func (s *darwinService) Stop() error {
	if err := exec.Command("launchctl", "stop", ServiceName).Run(); err != nil {
		return fmt.Errorf("failed to stop service: %w", err)
	}
	return nil
}
