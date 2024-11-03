package platform

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

type linuxService struct{}

const systemdServiceTemplate = `[Unit]
Description=%s

[Service]
ExecStart=%s -run
Restart=always
User=root
WorkingDirectory=%s

[Install]
WantedBy=multi-user.target
`

func (s *linuxService) Install(execPath string) error {
	installDir := GetInstallDir()
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return fmt.Errorf("failed to create installation directory: %w", err)
	}

	installedBinary := filepath.Join(installDir, "bin", filepath.Base(execPath))
	if err := os.MkdirAll(filepath.Dir(installedBinary), 0755); err != nil {
		return fmt.Errorf("failed to create bin directory: %w", err)
	}

	if err := copyFile(execPath, installedBinary); err != nil {
		return fmt.Errorf("failed to copy binary: %w", err)
	}

	servicePath := filepath.Join("/etc/systemd/system", ServiceName+".service")
	content := fmt.Sprintf(systemdServiceTemplate, ServiceDesc, installedBinary, installDir)

	if err := os.WriteFile(servicePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write service file: %w", err)
	}

	commands := [][]string{
		{"systemctl", "daemon-reload"},
		{"systemctl", "enable", ServiceName},
		{"systemctl", "start", ServiceName},
	}

	for _, args := range commands {
		if err := exec.Command(args[0], args[1:]...).Run(); err != nil {
			return fmt.Errorf("failed to execute %s: %w", args[0], err)
		}
	}
	return nil
}

func (s *linuxService) Uninstall() error {
	_ = exec.Command("systemctl", "stop", ServiceName).Run()
	_ = exec.Command("systemctl", "disable", ServiceName).Run()

	servicePath := filepath.Join("/etc/systemd/system", ServiceName+".service")
	if err := os.Remove(servicePath); err != nil {
		return fmt.Errorf("failed to remove service file: %w", err)
	}
	return nil
}

func (s *linuxService) Status() (bool, error) {
	output, err := exec.Command("systemctl", "is-active", ServiceName).Output()
	if err != nil {
		return false, nil
	}
	return string(output) == "active\n", nil
}

func (s *linuxService) Start() error {
	if err := exec.Command("systemctl", "start", ServiceName).Run(); err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}
	return nil
}

func (s *linuxService) Stop() error {
	if err := exec.Command("systemctl", "stop", ServiceName).Run(); err != nil {
		return fmt.Errorf("failed to stop service: %w", err)
	}
	return nil
}
