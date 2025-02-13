# goservice

This project is a fork of [Anshuman's go-service](https://github.com/ansxuman/go-service) and provides a library for managing system services across different platforms (Windows, Linux, and macOS). It allows developers to easily install, start, stop, and uninstall services from their Go applications.

This library provides a clean, modular architecture for managing system services that can run across different platforms.

## Features

- **Cross-Platform Service Management**: Supports Windows, Linux, and macOS.
- **Simple API**: Easy-to-use functions for service installation, uninstallation, starting, and stopping.
- **No External Dependencies**: Relies only on the Go standard library and platform-specific APIs.

## Getting Started

To use this library in your Go project:

```bash
go get github.com/benbenbenbenbenben/goservice
```

## Usage

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/benbenbenbenbenben/goservice"
)

func main() {
	// Define service parameters
	executablePath := "/path/to/your/executable" // or os.Executable() to self-install
	serviceName := "YourServiceName"
	serviceDescription := "Your Service Description"

	// Create a new service
	svc, err := goservice.NewService(executablePath, serviceName, serviceDescription)
	if err != nil {
		log.Fatalf("Failed to create service: %v", err)
	}

	// Install the service
	err = svc.Install()
	if err != nil {
		log.Fatalf("Failed to install service: %v", err)
	}
	defer func() {
		err := svc.Uninstall()
		if err != nil {
			log.Printf("Failed to uninstall service: %v", err)
		}
	}()

	// Start the service
	err = svc.Start()
	if err != nil {
		log.Fatalf("Failed to start service: %v", err)
	}
	defer func() {
		err := svc.Stop()
		if err != nil {
			log.Printf("Failed to stop service: %v", err)
		}
	}()
}
```

## Project Structure

```
goservice
├── service.go             # Main service implementation
├── example/               # Example service directory
│   ├── main.go            # Example service entry point
│   └── Taskfile.yml       # Taskfile for simplifying build and run processes
└── platform/              # Platform-specific implementations
    ├── copyFile.go        # Utility function for copying files
    ├── darwin.go          # macOS-specific service implementation
    ├── linux.go           # Linux-specific service implementation
    └── windows.go         # Windows-specific service implementation
```

## Example Service

To build and run the example service, follow the instructions in the `example/README.md` file. The example uses [Taskfile](https://taskfile.dev/) to simplify the build and run process.

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues to suggest improvements or report bugs.

## Acknowledgements

This project is based on the excellent work of [Anshuman](https://github.com/ansxuman) in creating the original `go-service` project. We would like to thank them for their contributions to the Go community.

## License

This project is licensed under the [MIT License](LICENSE).
