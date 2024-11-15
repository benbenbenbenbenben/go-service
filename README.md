# go-service

This project provides a clean, modular architecture for creating system services that can run across different platforms. It handles platform-specific service management while maintaining a consistent interface for service operations.

![GO-Service](https://github.com/user-attachments/assets/ccff528d-9897-4cd6-9e89-694feb11ad7c)


## Features

- **Clean Architecture**: Modular design with platform-agnostic core
- **Cross-Platform**: Supports Windows, Linux and macOS
- **Service Management**: Easy installation, status checking and lifecycle management 
- **Build Automation**: TaskFile for consistent build and management commands

## Project Overview

```bash
go-service/
├── Makefile                 # Build and installation automation
├── cmd/
│   └── service/
│       └── main.go          # Main entry point with CLI flags and command handling
├── internal/
│   ├── service/
│   │   └── service.go       # Core service implementation
│   └── platform/            # Platform-specific implementations
│       ├── config.go        # Configuration constants
│       ├── service.go       # Cross-platform service interface
│       ├── windows.go       # Windows-specific service management
│       ├── linux.go         # Linux-specific systemd service management
│       └── darwin.go        # macOS-specific launchd service management
└── go.mod                   # Go module definition
```

For detailed instructions, check the [Medium](https://medium.com/@ansxuman/building-cross-platform-system-services-in-go-a-step-by-step-guide-5784f96098b4) article for a step-by-step guide.If you don’t have a Medium membership, you can also find the article on [dev.to.](https://dev.to/ansxuman/building-cross-platform-system-services-in-go-a-step-by-step-guide-18mc)

## Donations

If you find this content useful, please consider donating to support its development and future improvements.

<a href="https://buymeacoffee.com/ansxuman" target="_blank"><img src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" alt="Buy Me A Coffee" style="height: 60px !important;width: 217px !important;" ></a>

