package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"go-service/internal/platform"
	"go-service/internal/service"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	install := flag.Bool("install", false, "Install the service")
	uninstall := flag.Bool("uninstall", false, "Uninstall the service")
	status := flag.Bool("status", false, "Check service status")
	start := flag.Bool("start", false, "Start the service")
	stop := flag.Bool("stop", false, "Stop the service")
	runWorker := flag.Bool("run", false, "Run the service worker")
	flag.Parse()

	if err := handleCommand(*install, *uninstall, *status, *start, *stop, *runWorker); err != nil {
		log.Fatal(err)
	}
}

func handleCommand(install, uninstall, status, start, stop, runWorker bool) error {
	platformSvc, err := platform.NewService()
	if err != nil {
		return err
	}

	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	switch {
	case install:
		return platformSvc.Install(execPath)
	case uninstall:
		return platformSvc.Uninstall()
	case status:
		running, err := platformSvc.Status()
		if err != nil {
			return err
		}
		fmt.Printf("Service is %s\n", map[bool]string{true: "running", false: "stopped"}[running])
		return nil
	case start:
		return platformSvc.Start()
	case stop:
		return platformSvc.Stop()
	case runWorker:
		return runService()
	default:
		return fmt.Errorf("no command specified")
	}
}

func runService() error {
	svc, err := service.New()
	if err != nil {
		return fmt.Errorf("failed to create service: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	log.Println("Starting service...")
	if err := svc.Start(ctx); err != nil {
		return fmt.Errorf("failed to start service: %w", err)
	}

	log.Println("Service started, waiting for shutdown signal...")
	<-sigChan
	log.Println("Shutdown signal received, stopping service...")

	if err := svc.Stop(); err != nil {
		return fmt.Errorf("failed to stop service: %w", err)
	}
	log.Println("Service stopped successfully")
	return nil
}
