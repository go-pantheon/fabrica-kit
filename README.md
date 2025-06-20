<div align="center">
  <h1>🏛️ FABRICA KIT</h1>
  <p><em>Essential engineered toolkit library for the go-pantheon ecosystem</em></p>
</div>

<p align="center">
<a href="https://github.com/go-pantheon/fabrica-kit/actions/workflows/test.yml"><img src="https://github.com/go-pantheon/fabrica-kit/workflows/Test/badge.svg" alt="Test Status"></a>
<a href="https://github.com/go-pantheon/fabrica-kit/releases"><img src="https://img.shields.io/github/v/release/go-pantheon/fabrica-kit" alt="Latest Release"></a>
<a href="https://pkg.go.dev/github.com/go-pantheon/fabrica-kit"><img src="https://pkg.go.dev/badge/github.com/go-pantheon/fabrica-kit" alt="GoDoc"></a>
<a href="https://goreportcard.com/report/github.com/go-pantheon/fabrica-kit"><img src="https://goreportcard.com/badge/github.com/go-pantheon/fabrica-kit" alt="Go Report Card"></a>
<a href="https://github.com/go-pantheon/fabrica-kit/blob/main/LICENSE"><img src="https://img.shields.io/github/license/go-pantheon/fabrica-kit" alt="License"></a>
<a href="https://deepwiki.com/go-pantheon/fabrica-kit"><img src="https://deepwiki.com/badge.svg" alt="Ask DeepWiki"></a>
</p>

> **Language**: [English](README.md) | [中文](README-zh.md)

## About Fabrica Kit

Fabrica Kit is a core toolkit library for the go-pantheon ecosystem, providing essential functionalities and integrations for building robust game server microservices. This kit builds upon the go-pantheon infrastructure to offer standardized components for logging, error handling, tracing, routing, and more, ensuring consistency and reliability across all go-pantheon services.

For more information, please check out: [deepwiki/go-pantheon/fabrica-kit](https://deepwiki.com/go-pantheon/fabrica-kit)

## About go-pantheon Ecosystem

**go-pantheon** is a high-performance, highly available game server cluster solution framework providing out-of-the-box game server infrastructure based on microservices architecture using [go-kratos](https://github.com/go-kratos/kratos). Fabrica Kit serves as the foundational toolkit that supports the core components:

- **Roma**: Game core logic services
- **Janus**: Gateway service for client connection handling and request forwarding
- **Lares**: Account service for user authentication and account management
- **Senate**: Backend management service providing operational interfaces

### Core Features

- 📝 **Structured Logging**: Advanced logging system with multi-format support, levels, and tracing integration
- 🛡️ **Error Handling**: Comprehensive API error standardization with predefined error types and handling
- 🔍 **Distributed Tracing**: OpenTelemetry-based tracing for microservices observability
- 📊 **Metrics Collection**: Built-in metrics collection and monitoring for performance tracking
- 🌐 **Routing & Load Balancing**: Intelligent routing with connection management and load balancing
- 📈 **Service Profiling**: Environment-aware configuration management and service metadata
- 🔧 **Context Extensions**: Enhanced context handling for request processing
- 🌍 **Network Utilities**: IP address handling and network-related tools
- 🔢 **Version Management**: Version control and compatibility checking utilities

## Toolkit Packages

### Structured Logging (`xlog/`)
Advanced logging system with comprehensive features:
- **Multi-format Support**: JSON and console logging formats
- **Level Management**: Debug, Info, Warn, Error levels with configurable thresholds
- **Tracing Integration**: Automatic trace ID and span ID injection
- **Metadata Enrichment**: Service name, version, profile, and node information

### Error Handling (`xerrors/`)
Comprehensive error management system:
- **API Error Standards**: Predefined error types for common scenarios
- **HTTP Status Mapping**: Automatic mapping to appropriate HTTP status codes
- **Error Context**: Rich error context with formatting support
- **Kratos Integration**: Seamless integration with go-kratos error handling

### Distributed Tracing (`trace/`)
OpenTelemetry-based tracing infrastructure:
- **OTLP Exporter**: HTTP-based trace exporter configuration
- **Service Identification**: Automatic service name and metadata tagging
- **Sampling Control**: Configurable sampling strategies
- **Multi-backend Support**: GORM, PostgreSQL, Redis tracing instrumentation

### Metrics Collection (`metrics/`)
Performance monitoring and metrics:
- **Request Metrics**: Automatic request counting and duration tracking
- **OpenTelemetry Integration**: Standard metrics collection using OTEL
- **Server & Client Middleware**: Bidirectional metrics collection
- **Custom Metrics**: Support for application-specific metrics

### Routing System (`router/`)
Intelligent routing and load balancing:
- **Load Balancing**: Multiple balancing algorithms and strategies
- **Connection Management**: Efficient connection pooling and lifecycle management
- **Route Tables**: Dynamic route configuration and management
- **Service Discovery**: Integration with service discovery mechanisms

### Other Components
- **Profile** (`profile/`): Environment-aware configuration and service metadata
- **Context Extensions** (`xcontext/`): Enhanced context handling utilities
- **IP Utilities** (`ip/`): Network address processing and validation
- **Version Tools** (`version/`): Version management and compatibility checking

## Technology Stack

| Technology/Component | Purpose                             | Version  |
| -------------------- | ----------------------------------- | -------- |
| Go                   | Primary development language        | 1.24+    |
| go-kratos            | Microservices framework             | v2.8.4   |
| OpenTelemetry        | Distributed tracing and metrics     | v1.35.0  |
| Zap                  | High-performance structured logging | v1.27.0  |
| go-redis             | Redis client                        | v9.7.3   |
| gRPC                 | Remote procedure call               | v1.71.1  |
| GORM                 | Object-relational mapping           | v1.25.12 |

## Requirements

- Go 1.24+

## Quick Start

### Installation

```bash
go get github.com/go-pantheon/fabrica-kit
```

### Initialize Development Environment

```bash
make init
```

### Run Tests

```bash
make test
```

## Usage Examples

### Structured Logging with Tracing

```go
package main

import (
    "github.com/go-pantheon/fabrica-kit/xlog"
    "github.com/go-pantheon/fabrica-kit/trace"
)

func main() {
    // Initialize tracing
    err := trace.Init("http://localhost:4318/v1/traces", "my-service", "dev", "blue")
    if err != nil {
        panic(err)
    }

    // Initialize logger with comprehensive metadata
    logger := xlog.Init("zap", "info", "dev", "blue", "my-service", "v1.0.0", "node-1")

    // Logging with automatic trace context
    logger.Info("Service started successfully")
    logger.Error("Database connection failed", "error", err)
}
```

### API Error Handling

```go
package main

import (
    "fmt"
    "github.com/go-pantheon/fabrica-kit/xerrors"
)

func validateUser(userID string) error {
    if userID == "" {
        return xerrors.APIParamInvalid("user ID cannot be empty")
    }

    if len(userID) < 3 {
        return xerrors.APIParamInvalid("user ID must be at least %d characters", 3)
    }

    // Check if user exists
    if !userExists(userID) {
        return xerrors.APINotFound("user %s not found", userID)
    }

    return nil
}

func handleUserRequest(userID string) {
    if err := validateUser(userID); err != nil {
        // Error is automatically converted to appropriate HTTP status
        fmt.Printf("Request failed: %v\n", err)
    }
}
```

### Metrics Collection

```go
package main

import (
    "github.com/go-kratos/kratos/v2"
    "github.com/go-kratos/kratos/v2/transport/grpc"
    "github.com/go-pantheon/fabrica-kit/metrics"
)

func main() {
    // Initialize metrics system
    metrics.Init("my-service")

    // Create gRPC server with metrics middleware
    server := grpc.NewServer(
        grpc.Middleware(
            metrics.Server(), // Automatic request/duration metrics
        ),
    )

    app := kratos.New(
        kratos.Name("my-service"),
        kratos.Server(server),
    )

    if err := app.Run(); err != nil {
        panic(err)
    }
}
```

### Environment-aware Configuration

```go
package main

import (
    "fmt"
    "github.com/go-pantheon/fabrica-kit/profile"
)

func main() {
    // Environment-aware behavior
    if profile.IsDev() {
        fmt.Println("Running in development mode")
        // Enable debug features
    }

    // Configure based on environment
    var logLevel string
    switch {
    case profile.IsDev():
        logLevel = "debug"
    case profile.IsTestStr("test"):
        logLevel = "info"
    default:
        logLevel = "warn"
    }

    fmt.Printf("Log level set to: %s\n", logLevel)
}
```

## Project Structure

```
.
├── xlog/               # Structured logging system
│   ├── log.go          # Main logging functionality
│   └── gorm.go         # GORM integration
├── xerrors/            # Error handling framework
│   ├── apierrors.go    # API error definitions
│   ├── kiterrors.go    # Kit-specific errors
│   └── logouterrors.go # Logout error handlers
├── trace/              # Distributed tracing
│   ├── trace.go        # Core tracing functionality
│   ├── gorm/           # GORM tracing instrumentation
│   ├── postgresql/     # PostgreSQL tracing
│   └── redis/          # Redis tracing
├── metrics/            # Metrics collection
│   ├── metrics.go      # Core metrics functionality
│   ├── postgresql/     # PostgreSQL metrics
│   └── redis/          # Redis metrics
├── router/             # Routing and load balancing
│   ├── constants.go    # Router constants
│   ├── balancer/       # Load balancing algorithms
│   ├── conn/           # Connection management
│   └── routetable/     # Routing table management
├── profile/            # Service profiling and metadata
├── xcontext/           # Context extensions
├── ip/                 # IP address utilities
└── version/            # Version management
```

## Integration with go-pantheon Components

Fabrica Kit is designed to be imported by other go-pantheon components to provide common functionality:

```go
import (
    // Structured logging for all services
    "github.com/go-pantheon/fabrica-kit/xlog"

    // Routing and load balancing in Janus
    "github.com/go-pantheon/fabrica-kit/router"

    // Error handling in Lares and Roma
    "github.com/go-pantheon/fabrica-kit/xerrors"

    // Distributed tracing across all services
    "github.com/go-pantheon/fabrica-kit/trace"

    // Metrics collection for monitoring
    "github.com/go-pantheon/fabrica-kit/metrics"

    // Environment-aware configuration
    "github.com/go-pantheon/fabrica-kit/profile"
)
```

## Development Guide

### License Compliance

The project enforces license compliance for all dependencies. We only allow the following licenses:
- MIT
- Apache-2.0
- BSD-2-Clause
- BSD-3-Clause
- ISC
- MPL-2.0

License checks are performed:
- Automatically in CI/CD pipelines
- Locally via pre-commit hooks
- Manually using `make license-check`

### Testing

Run the complete test suite:

```bash
# Run all tests with coverage
make test

# Run linting
make lint

# Run go vet
make vet
```

### Adding New Features

When adding new toolkit components:

1. Create a new package or add to an existing one based on functionality
2. Implement features with proper error handling using `xerrors`
3. Add comprehensive logging using `xlog`
4. Include distributed tracing support where applicable
5. Write comprehensive unit tests with edge cases covered
6. Ensure thread safety where applicable
7. Document usage with clear examples
8. Run tests: `make test`
9. Update documentation for any API changes

### Middleware Integration

When creating new middleware:

```go
func MyMiddleware() middleware.Middleware {
    return func(handler middleware.Handler) middleware.Handler {
        return func(ctx context.Context, req interface{}) (interface{}, error) {
            // Use fabrica-kit components
            logger := xlog.FromContext(ctx)
            logger.Info("Processing request")

            reply, err := handler(ctx, req)
            if err != nil {
                // Standardized error handling
                return nil, xerrors.APIInternalError("request failed: %v", err)
            }

            return reply, nil
        }
    }
}
```

### Contribution Guidelines

1. Fork this repository
2. Create a feature branch from `main`
3. Implement changes with comprehensive tests
4. Ensure all tests pass and linting is clean
5. Follow the established patterns for error handling and logging
6. Update documentation for any API changes
7. Submit a Pull Request with clear description

## Performance Considerations

- **Logging**: Structured logging is optimized for high-throughput scenarios
- **Tracing**: Sampling strategies can be configured to balance observability and performance
- **Metrics**: Metrics collection uses atomic operations and lock-free data structures
- **Error Handling**: Error creation is lightweight with minimal allocations
- **Routing**: Connection pooling and load balancing algorithms are optimized for low latency
- **Memory Usage**: All components are designed to minimize memory allocations and GC pressure

## License

This project is licensed under the terms specified in the LICENSE file.
