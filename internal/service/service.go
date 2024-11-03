package service

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"time"

	"go-service/internal/platform"
)

type Service struct {
	logFile string
	stop    chan struct{}
	wg      sync.WaitGroup
	started bool
	mu      sync.Mutex
}

func New() (*Service, error) {
	installDir := platform.GetInstallDir()
	if installDir == "" {
		return nil, fmt.Errorf("unsupported operating system: %s", runtime.GOOS)
	}

	logFile := filepath.Join(installDir, "logs", platform.LogFileName)

	return &Service{
		logFile: logFile,
		stop:    make(chan struct{}),
	}, nil
}

func (s *Service) Start(ctx context.Context) error {
	s.mu.Lock()
	if s.started {
		s.mu.Unlock()
		return fmt.Errorf("service already started")
	}
	s.started = true
	s.mu.Unlock()

	if err := os.MkdirAll(filepath.Dir(s.logFile), 0755); err != nil {
		return fmt.Errorf("failed to create log directory: %w", err)
	}

	s.wg.Add(1)
	go s.run(ctx)

	return nil
}

func (s *Service) Stop() error {
	s.mu.Lock()
	if !s.started {
		s.mu.Unlock()
		return fmt.Errorf("service not started")
	}
	s.mu.Unlock()

	close(s.stop)
	s.wg.Wait()

	s.mu.Lock()
	s.started = false
	s.mu.Unlock()

	return nil
}

func (s *Service) run(ctx context.Context) {
	defer s.wg.Done()
	log.Printf("Service started, logging to: %s\n", s.logFile)

	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	if err := s.writeLog(); err != nil {
		log.Printf("Error writing initial log: %v\n", err)
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("Service stopping due to context cancellation")
			return
		case <-s.stop:
			log.Println("Service stopping due to stop signal")
			return
		case <-ticker.C:
			if err := s.writeLog(); err != nil {
				log.Printf("Error writing log: %v\n", err)
			}
		}
	}
}

func (s *Service) writeLog() error {
	f, err := os.OpenFile(s.logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open log file: %w", err)
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("[%s] Hello World\n", time.Now().Format(time.RFC3339)))
	if err != nil {
		return fmt.Errorf("failed to write to log file: %w", err)
	}
	return nil
}
