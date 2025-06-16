<p align="center">
<h1 class="logo">FABRICA KIT</h1>
</p>

<p align="center">
<a href="https://github.com/go-pantheon/fabrica-kit/actions"><img src="https://github.com/go-pantheon/fabrica-kit/workflows/Go/badge.svg" alt="Build Status"></a>
<a href="https://pkg.go.dev/github.com/go-pantheon/fabrica-kit"><img src="https://pkg.go.dev/badge/github.com/go-pantheon/fabrica-kit/v2" alt="GoDoc"></a>
<a href="https://codecov.io/gh/go-pantheon/fabrica-kit"><img src="https://codecov.io/gh/go-pantheon/fabrica-kit/master/graph/badge.svg" alt="codeCov"></a>
<a href="https://goreportcard.com/report/github.com/go-pantheon/fabrica-kit"><img src="https://goreportcard.com/badge/github.com/go-pantheon/fabrica-kit" alt="Go Report Card"></a>
<a href="https://github.com/go-pantheon/fabrica-kit/blob/main/LICENSE"><img src="https://img.shields.io/github/license/go-pantheon/fabrica-kit" alt="License"></a>
<a href="https://deepwiki.com/go-pantheon/fabrica-kit"><img src="https://deepwiki.com/badge.svg" alt="Ask DeepWiki"></a>
</p>

## About Fabrica Kit

Fabrica Kit is a core toolkit library for the go-pantheon ecosystem, providing essential functionalities and integrations for building robust game server microservices. This kit builds upon the go-pantheon infrastructure to offer standardized components for logging, error handling, tracing, routing, and more.

For more, please check out: [deepwiki/go-pantheon/fabrica-kit](https://deepwiki.com/go-pantheon/fabrica-kit)

## About go-pantheon Ecosystem

**go-pantheon** is a high-performance, highly available game server cluster solution framework based on [go-kratos](https://github.com/go-kratos/kratos). Fabrica Kit, as a framework toolkit, provides support for the following core components:

- **Roma**: Game core logic services
- **Janus**: Gateway service for client connection handling and request forwarding
- **Lares**: Account service for user authentication and account management
- **Senate**: Backend management service providing operational interfaces

## Core Features

- 📝 Standardized logging system (xlog/)
- 🔍 Tracing and metrics collection (trace/, metrics/)
- 🌐 Routing and load balancing (router/)
- 🛡️ Error handling and API error standardization (xerrors/)
- 🔄 Communication tunnel abstraction (tunnel/)
- 📊 Context extensions (xcontext/)
- 📈 Project metadata (profile/)
- 🌍 IP address handling tools (ip/)
- 🔢 Version control and compatibility checking (version/)

## Technology Stack

| Technology/Component | Purpose                             | Version  |
| -------------------- | ----------------------------------- | -------- |
| Go                   | Primary development language        | 1.23+    |
| go-kratos            | Microservices framework             | v2.8.4   |
| OpenTelemetry        | Distributed tracing and metrics     | v1.35.0  |
| Zap                  | High-performance structured logging | v1.27.0  |
| go-redis             | Redis client                        | v9.7.3   |
| gRPC                 | Remote procedure call               | v1.71.1  |
| GORM                 | Object-relational mapping           | v1.25.12 |

## Quick Start

### Installation

```bash
go get github.com/go-pantheon/fabrica-kit
```

## Project Structure

```
.
├── xlog/               # Logging tools
├── xerrors/            # Error handling
├── trace/              # Distributed tracing
├── metrics/            # Metrics collection
├── router/             # Routing and load balancing
│   ├── balancer/       # Load balancing implementation
│   ├── conn/           # Connection management
│   └── routetable/     # Routing table
├── tunnel/             # Communication tunnel abstraction
├── xcontext/           # Context extensions
├── profile/            # Project metadata
├── ip/                 # IP address handling
└── version/            # Version tools
```

## Integration with go-pantheon Components

Fabrica Kit is designed to be imported by other go-pantheon components to provide common functionality:

```go
import (
    // For logging in Roma
    "github.com/go-pantheon/fabrica-kit/xlog"

    // For routing in Janus
    "github.com/go-pantheon/fabrica-kit/router"

    // For error handling in Lares
    "github.com/go-pantheon/fabrica-kit/xerrors"

    // For tracing in all services
    "github.com/go-pantheon/fabrica-kit/trace"
)
```

## Development Guide

### Adding New Features

When adding new features:

1. Create a new package or add to an existing one based on functionality
2. Implement features with proper error handling
3. Write comprehensive unit tests
4. Document usage with examples
5. Run tests to ensure functionality works correctly

### Contribution Guidelines

1. Fork this repository
2. Create a feature branch
3. Submit changes with comprehensive tests
4. Ensure all tests pass
5. Submit a Pull Request

## License

This project is licensed under the terms specified in the LICENSE file.
