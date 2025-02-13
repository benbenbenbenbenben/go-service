package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/benbenbenbenbenben/goservice"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)

	install := flag.Bool("install", false, "Install the service")
	uninstall := flag.Bool("uninstall", false, "Uninstall the service")
	startService := flag.Bool("start", false, "Start the service")
	stopService := flag.Bool("stop", false, "Stop the service")
	statusService := flag.Bool("status", false, "Get the service status")
	flag.Parse()

	executablePath, err := os.Executable()
	if err != nil {
		log.Fatalf("Failed to get executable path: %v", err)
	}

	svc, err := goservice.NewService(executablePath, "HelloWorldService", "A simple Hello World service")
	if err != nil {
		log.Fatalf("Failed to create service: %v", err)
	}

	switch {
	case *install:
		if err := svc.Install(); err != nil {
			log.Fatalf("Failed to install service: %v", err)
		}
		fmt.Println("Service installed successfully.")
		return
	case *uninstall:
		if err := svc.Uninstall(); err != nil {
			log.Fatalf("Failed to uninstall service: %v", err)
		}
		fmt.Println("Service uninstalled successfully.")
		return
	case *startService:
		if err := svc.Start(); err != nil {
			log.Fatalf("Failed to start service: %v", err)
		}
		fmt.Println("Service started successfully.")
		return
	case *stopService:
		if err := svc.Stop(); err != nil {
			log.Fatalf("Failed to stop service: %v", err)
		}
		fmt.Println("Service stopped successfully.")
		return
	case *statusService:
		status, err := svc.Status()
		if err != nil {
			log.Fatalf("Failed to get service status: %v", err)
		}
		if status {
			fmt.Println("Service is running.")
		} else {
			fmt.Println("Service is not running.")
		}
		return
	default:
		Run()
		return
	}
}

func Run() {
	// Print hello world every 5 seconds until the program is terminated
	for {
		fmt.Println("Hello World!")
		time.Sleep(5 * time.Second)
	}
}
