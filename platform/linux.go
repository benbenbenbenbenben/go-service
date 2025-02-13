package platform

import (
	"fmt"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

type linuxService struct{
	executablePath string
	serviceDisplay string
	serviceDesc string
}

const systemdServiceTemplate = `[Unit]
Description=%s

[Service]
ExecStart=%s
Restart=always
User=root
WorkingDirectory=%s

[Install]
WantedBy=multi-user.target
`

func (s *linuxService) Install() error {
	installDir := s.GetInstallDir()
	if err := os.MkdirAll(installDir, 0755); err != nil {
		return fmt.Errorf("failed to create installation directory: %w", err)
	}

	installedBinary := filepath.Join(installDir, "bin", filepath.Base(s.executablePath))
	if err := os.MkdirAll(filepath.Dir(installedBinary), 0755); err != nil {
		return fmt.Errorf("failed to create bin directory: %w", err)
	}

	if err := copyFile(s.executablePath, installedBinary); err != nil {
		return fmt.Errorf("failed to copy binary: %w", err)
	}

	servicePath := filepath.Join("/etc/systemd/system", s.ServiceName()+".service")
	content := fmt.Sprintf(systemdServiceTemplate, s.serviceDesc, installedBinary, installDir)

	if err := os.WriteFile(servicePath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write service file: %w", err)
	}

	commands := [][]string{
		{"systemctl", "daemon-reload"},
		{"systemctl", "enable", s.ServiceName()},
		{"systemctl", "start", s.ServiceName()},
	}

	for _, args := range commands {
		if err := exec.Command(args[0], args[1:]...).Run(); err != nil {
			return fmt.Errorf("failed to execute %s: %w", args[0], err)
		}
	}
	return nil
}

func (s *linuxService) Uninstall() error {
	_ = exec.Command("systemctl", "stop", s.ServiceName()).Run()
	_ = exec.Command("systemctl", "disable", s.ServiceName()).Run()

	servicePath := filepath.Join("/etc/systemd/system", s.ServiceName()+".service")
	if err := os.Remove(servicePath); err != nil {
		return fmt.Errorf("failed to remove service file: %w", err)
	}
	return nil
}

func (s *linuxService) Status() (bool, error) {
	output, err := exec.Command("systemctl", "is-active", s.ServiceName()).Output()
	if err != nil {
		return false, nil
	}
	return string(output) == "active\n", nil
}

func (s *linuxService) Start() error {
	if err := exec.Command("systemctl", "start", s.ServiceName()).Run(); err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}
	return nil
}

func (s *linuxService) Stop() error {
	if err := exec.Command("systemctl", "stop", s.ServiceName()).Run(); err != nil {
		return fmt.Errorf("failed to stop service: %w", err)
	}
	return nil
}

func (s *linuxService) GetInstallDir() string {
	return "/opt/" + s.ServiceName()
}

func (s *linuxService) ServiceName() string {
	return path.Base(s.executablePath)
}

func (s *linuxService) ServiceDisplayName() string {
	return s.serviceDisplay
}

func (s *linuxService) ServiceDescription() string {
	return s.serviceDesc
}