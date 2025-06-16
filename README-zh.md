<div align="center">
  <h1>🏛️ FABRICA KIT</h1>
  <p><em>go-pantheon 生态系统的核心工程化工具包</em></p>
</div>

<p align="center">
<a href="https://github.com/go-pantheon/fabrica-kit/actions/workflows/test.yml"><img src="https://github.com/go-pantheon/fabrica-kit/workflows/Test/badge.svg" alt="Test Status"></a>
<a href="https://github.com/go-pantheon/fabrica-kit/releases"><img src="https://img.shields.io/github/v/release/go-pantheon/fabrica-kit" alt="Latest Release"></a>
<a href="https://pkg.go.dev/github.com/go-pantheon/fabrica-kit"><img src="https://pkg.go.dev/badge/github.com/go-pantheon/fabrica-kit" alt="GoDoc"></a>
<a href="https://goreportcard.com/report/github.com/go-pantheon/fabrica-kit"><img src="https://goreportcard.com/badge/github.com/go-pantheon/fabrica-kit" alt="Go Report Card"></a>
<a href="https://github.com/go-pantheon/fabrica-kit/blob/main/LICENSE"><img src="https://img.shields.io/github/license/go-pantheon/fabrica-kit" alt="License"></a>
<a href="https://deepwiki.com/go-pantheon/fabrica-kit"><img src="https://deepwiki.com/badge.svg" alt="Ask DeepWiki"></a>
</p>

> **语言**: [English](README.md) | [中文](README-zh.md)

## 关于 Fabrica Kit

Fabrica Kit 是 go-pantheon 生态系统的核心工具包库，为构建健壮的游戏服务器微服务提供基础功能和集成。该工具包基于 go-pantheon 基础设施构建，提供标准化的日志记录、错误处理、追踪、路由等组件，确保所有 go-pantheon 服务的一致性和可靠性。

更多信息请查看：[deepwiki/go-pantheon/fabrica-kit](https://deepwiki.com/go-pantheon/fabrica-kit)

## 关于 go-pantheon 生态系统

**go-pantheon** 是一个高性能、高可用的游戏服务器集群解决方案框架，基于微服务架构使用 [go-kratos](https://github.com/go-kratos/kratos) 提供开箱即用的游戏服务器基础设施。Fabrica Kit 作为基础工具包，支持以下核心组件：

- **Roma**: 游戏核心逻辑服务
- **Janus**: 网关服务，处理客户端连接和请求转发
- **Lares**: 账户服务，用于用户认证和账户管理
- **Senate**: 后台管理服务，提供运营接口

### 核心特性

- 📝 **结构化日志**: 具有多格式支持、级别管理和追踪集成的高级日志系统
- 🛡️ **错误处理**: 具有预定义错误类型和处理机制的全面API错误标准化
- 🔍 **分布式追踪**: 基于OpenTelemetry的微服务可观测性追踪
- 📊 **指标收集**: 内置指标收集和性能监控
- 🌐 **路由与负载均衡**: 具有连接管理和负载均衡的智能路由
- 📈 **服务配置**: 环境感知的配置管理和服务元数据
- 🔧 **上下文扩展**: 增强的请求处理上下文处理
- 🌍 **网络工具**: IP地址处理和网络相关工具
- 🔢 **版本管理**: 版本控制和兼容性检查工具

## 工具包组件

### 结构化日志 (`xlog/`)
具有全面功能的高级日志系统：
- **多格式支持**: JSON和控制台日志格式
- **级别管理**: Debug、Info、Warn、Error级别，可配置阈值
- **追踪集成**: 自动注入trace ID和span ID
- **元数据丰富**: 服务名、版本、配置文件和节点信息

### 错误处理 (`xerrors/`)
全面的错误管理系统：
- **API错误标准**: 常见场景的预定义错误类型
- **HTTP状态映射**: 自动映射到适当的HTTP状态码
- **错误上下文**: 支持格式化的丰富错误上下文
- **Kratos集成**: 与go-kratos错误处理的无缝集成

### 分布式追踪 (`trace/`)
基于OpenTelemetry的追踪基础设施：
- **OTLP导出器**: 基于HTTP的追踪导出器配置
- **服务识别**: 自动服务名和元数据标记
- **采样控制**: 可配置的采样策略
- **多后端支持**: GORM、PostgreSQL、Redis追踪仪表

### 指标收集 (`metrics/`)
性能监控和指标：
- **请求指标**: 自动请求计数和持续时间跟踪
- **OpenTelemetry集成**: 使用OTEL的标准指标收集
- **服务器和客户端中间件**: 双向指标收集
- **自定义指标**: 支持应用程序特定的指标

### 路由系统 (`router/`)
智能路由和负载均衡：
- **负载均衡**: 多种平衡算法和策略
- **连接管理**: 高效的连接池和生命周期管理
- **路由表**: 动态路由配置和管理
- **服务发现**: 与服务发现机制集成

### 其他组件
- **Profile** (`profile/`): 环境感知配置和服务元数据
- **上下文扩展** (`xcontext/`): 增强的上下文处理工具
- **IP工具** (`ip/`): 网络地址处理和验证
- **版本工具** (`version/`): 版本管理和兼容性检查

## 技术栈

| 技术/组件     | 用途             | 版本     |
| ------------- | ---------------- | -------- |
| Go            | 主要开发语言     | 1.23+    |
| go-kratos     | 微服务框架       | v2.8.4   |
| OpenTelemetry | 分布式追踪和指标 | v1.35.0  |
| Zap           | 高性能结构化日志 | v1.27.0  |
| go-redis      | Redis客户端      | v9.7.3   |
| gRPC          | 远程过程调用     | v1.71.1  |
| GORM          | 对象关系映射     | v1.25.12 |

## 系统要求

- Go 1.23+

## 快速开始

### 安装

```bash
go get github.com/go-pantheon/fabrica-kit
```

### 初始化开发环境

```bash
make init
```

### 运行测试

```bash
make test
```

## 使用示例

### 结构化日志与追踪

```go
package main

import (
    "github.com/go-pantheon/fabrica-kit/xlog"
    "github.com/go-pantheon/fabrica-kit/trace"
)

func main() {
    // 初始化追踪
    err := trace.Init("http://localhost:4318/v1/traces", "my-service", "dev", "blue")
    if err != nil {
        panic(err)
    }

    // 初始化带有全面元数据的日志器
    logger := xlog.Init("zap", "info", "dev", "blue", "my-service", "v1.0.0", "node-1")

    // 带有自动追踪上下文的日志记录
    logger.Info("服务启动成功")
    logger.Error("数据库连接失败", "error", err)
}
```

### API错误处理

```go
package main

import (
    "fmt"
    "github.com/go-pantheon/fabrica-kit/xerrors"
)

func validateUser(userID string) error {
    if userID == "" {
        return xerrors.APIParamInvalid("用户ID不能为空")
    }

    if len(userID) < 3 {
        return xerrors.APIParamInvalid("用户ID必须至少%d个字符", 3)
    }

    // 检查用户是否存在
    if !userExists(userID) {
        return xerrors.APINotFound("用户%s未找到", userID)
    }

    return nil
}

func handleUserRequest(userID string) {
    if err := validateUser(userID); err != nil {
        // 错误自动转换为适当的HTTP状态
        fmt.Printf("请求失败: %v\n", err)
    }
}
```

### 指标收集

```go
package main

import (
    "github.com/go-kratos/kratos/v2"
    "github.com/go-kratos/kratos/v2/transport/grpc"
    "github.com/go-pantheon/fabrica-kit/metrics"
)

func main() {
    // 初始化指标系统
    metrics.Init("my-service")

    // 创建带有指标中间件的gRPC服务器
    server := grpc.NewServer(
        grpc.Middleware(
            metrics.Server(), // 自动请求/持续时间指标
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

### 环境感知配置

```go
package main

import (
    "fmt"
    "github.com/go-pantheon/fabrica-kit/profile"
)

func main() {
    // 环境感知行为
    if profile.IsDev() {
        fmt.Println("在开发模式下运行")
        // 启用调试功能
    }

    // 基于环境配置
    var logLevel string
    switch {
    case profile.IsDev():
        logLevel = "debug"
    case profile.IsTestStr("test"):
        logLevel = "info"
    default:
        logLevel = "warn"
    }

    fmt.Printf("日志级别设置为: %s\n", logLevel)
}
```

## 项目结构

```
.
├── xlog/               # 结构化日志系统
│   ├── log.go          # 主日志功能
│   └── gorm.go         # GORM集成
├── xerrors/            # 错误处理框架
│   ├── apierrors.go    # API错误定义
│   ├── kiterrors.go    # Kit特定错误
│   └── logouterrors.go # 登出错误处理器
├── trace/              # 分布式追踪
│   ├── trace.go        # 核心追踪功能
│   ├── gorm/           # GORM追踪仪表
│   ├── postgresql/     # PostgreSQL追踪
│   └── redis/          # Redis追踪
├── metrics/            # 指标收集
│   ├── metrics.go      # 核心指标功能
│   ├── postgresql/     # PostgreSQL指标
│   └── redis/          # Redis指标
├── router/             # 路由和负载均衡
│   ├── constants.go    # 路由常量
│   ├── balancer/       # 负载均衡算法
│   ├── conn/           # 连接管理
│   └── routetable/     # 路由表管理
├── profile/            # 服务配置和元数据
├── xcontext/           # 上下文扩展
├── ip/                 # IP地址工具
└── version/            # 版本管理
```

## 与 go-pantheon 组件集成

Fabrica Kit 设计为供其他 go-pantheon 组件导入以提供通用功能：

```go
import (
    // 所有服务的结构化日志
    "github.com/go-pantheon/fabrica-kit/xlog"

    // Janus中的路由和负载均衡
    "github.com/go-pantheon/fabrica-kit/router"

    // Lares和Roma中的错误处理
    "github.com/go-pantheon/fabrica-kit/xerrors"

    // 所有服务的分布式追踪
    "github.com/go-pantheon/fabrica-kit/trace"

    // 监控的指标收集
    "github.com/go-pantheon/fabrica-kit/metrics"

    // 环境感知配置
    "github.com/go-pantheon/fabrica-kit/profile"
)
```

## 开发指南

### 许可证合规性

项目对所有依赖项强制执行许可证合规性。我们只允许以下许可证：
- MIT
- Apache-2.0
- BSD-2-Clause
- BSD-3-Clause
- ISC
- MPL-2.0

许可证检查执行：
- 在CI/CD管道中自动执行
- 通过pre-commit钩子在本地执行
- 使用`make license-check`手动执行

### 测试

运行完整的测试套件：

```bash
# 运行所有测试并生成覆盖率报告
make test

# 运行代码检查
make lint

# 运行go vet
make vet
```

### 添加新功能

添加新的工具包组件时：

1. 根据功能创建新包或添加到现有包
2. 使用`xerrors`实现具有适当错误处理的功能
3. 使用`xlog`添加全面的日志记录
4. 在适用的地方包含分布式追踪支持
5. 编写覆盖边界情况的全面单元测试
6. 确保在适用的地方线程安全
7. 用清晰的示例记录使用方法
8. 运行测试：`make test`
9. 更新任何API更改的文档

### 中间件集成

创建新中间件时：

```go
func MyMiddleware() middleware.Middleware {
    return func(handler middleware.Handler) middleware.Handler {
        return func(ctx context.Context, req interface{}) (interface{}, error) {
            // 使用fabrica-kit组件
            logger := xlog.FromContext(ctx)
            logger.Info("处理请求")

            reply, err := handler(ctx, req)
            if err != nil {
                // 标准化错误处理
                return nil, xerrors.APIInternalError("请求失败: %v", err)
            }

            return reply, nil
        }
    }
}
```

### 贡献指南

1. Fork此仓库
2. 从`main`创建功能分支
3. 实现带有全面测试的更改
4. 确保所有测试通过且代码检查干净
5. 遵循已建立的错误处理和日志记录模式
6. 更新任何API更改的文档
7. 提交带有清晰描述的Pull Request

## 性能考虑

- **日志记录**: 结构化日志针对高吞吐量场景进行优化
- **追踪**: 可以配置采样策略以平衡可观测性和性能
- **指标**: 指标收集使用原子操作和无锁数据结构
- **错误处理**: 错误创建轻量级，分配最少
- **路由**: 连接池和负载均衡算法针对低延迟优化
- **内存使用**: 所有组件都设计为最小化内存分配和GC压力

## 许可证

此项目根据LICENSE文件中指定的条款进行许可。
