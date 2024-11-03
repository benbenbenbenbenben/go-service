package platform

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type windowsService struct{}

func (s *windowsService) Install(execPath string) error {
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

	cmd := exec.Command("sc", "create", ServiceName,
		"binPath=", fmt.Sprintf("\"%s\" -run", installedBinary),
		"DisplayName=", ServiceDisplay,
		"start=", "auto",
		"obj=", "LocalSystem")

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	descCmd := exec.Command("sc", "description", ServiceName, ServiceDesc)
	if err := descCmd.Run(); err != nil {
		return fmt.Errorf("failed to set service description: %w", err)
	}

	if err := exec.Command("sc", "start", ServiceName).Run(); err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}
	return nil
}

func (s *windowsService) Uninstall() error {
	_ = exec.Command("sc", "stop", ServiceName).Run()
	if err := exec.Command("sc", "delete", ServiceName).Run(); err != nil {
		return fmt.Errorf("failed to delete service: %w", err)
	}

	// Clean up installation directory
	installDir := GetInstallDir()
	if err := os.RemoveAll(installDir); err != nil {
		return fmt.Errorf("failed to remove installation directory: %w", err)
	}
	return nil
}
func (s *windowsService) Status() (bool, error) {
	output, err := exec.Command("sc", "query", ServiceName).Output()
	if err != nil {
		return false, nil
	}
	return strings.Contains(string(output), "RUNNING"), nil
}

func (s *windowsService) Start() error {
	if err := exec.Command("sc", "start", ServiceName).Run(); err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}
	return nil
}

func (s *windowsService) Stop() error {
	if err := exec.Command("sc", "stop", ServiceName).Run(); err != nil {
		return fmt.Errorf("failed to stop service: %w", err)
	}
	return nil
}
