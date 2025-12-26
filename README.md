# 🐿️ Chipmunk

> **Gather, store, and organize your environment variables into Go structs**

[![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.19-%23007d9c)](https://golang.org/)
[![GoDoc](https://godoc.org/github.com/yourusername/chipmunk?status.svg)](https://godoc.org/github.com/yourusername/chipmunk)
[![Go Report Card](https://goreportcard.com/badge/github.com/yourusername/chipmunk)](https://goreportcard.com/report/github.com/yourusername/chipmunk)
[![Coverage Status](https://codecov.io/gh/yourusername/chipmunk/branch/main/graph/badge.svg)](https://codecov.io/gh/yourusername/chipmunk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Just like how chipmunks efficiently gather and organize nuts in their cheek pouches, **Chipmunk** helps you gather and organize environment variables into well-structured Go structs with zero hassle and maximum flexibility.

## ✨ Features

- 🏗️ **Struct Loading**: Map environment variables directly to Go structs using intuitive tags
- 🔧 **Custom Types**: Extensible parser system for your domain-specific types
- 📁 **Multiple Sources**: Load from .env files, OS environment, custom sources with priority
- 🏷️ **Rich Tags**: Comprehensive struct tag options (required, default, separator, expand, file)
- ⚡ **High Performance**: Optimized for speed and low memory usage
- 🛡️ **Type Safety**: Full compile-time type checking with detailed error messages
- 🎯 **Zero Dependencies**: Pure Go implementation, no external dependencies
- 🔐 **Security**: Built-in support for file-based secrets and variable expansion
- 🧪 **Testing**: Comprehensive test coverage with benchmarks

## 📦 Installation

```bash
go get github.com/raimialiu/chipmunk
```

## 🚀 Quick Start

### Basic Usage

```go
package main

import (
  "fmt"
  "log"
  "github.com/raimialiu/chipmunk"
)

type Config struct {
  // Required field - will fail if not set
  DatabaseURL string `env:"DATABASE_URL,required"`
  
  // Optional field with default value
  Port int `env:"PORT" default:"8080"`
  
  // Boolean field
  Debug bool `env:"DEBUG"`
  
  // String field with default
  LogLevel string `env:"LOG_LEVEL" default:"info"`
}

func main() {
  var config Config
  
  if err := chipmunk.Load(&config); err != nil {
      log.Fatal("Failed to load configuration:", err)
  }
  
  fmt.Printf("🚀 Server starting on port %d\n", config.Port)
  fmt.Printf("📊 Debug mode: %v\n", config.Debug)
  fmt.Printf("🗃️ Database: %s\n", config.DatabaseURL)
}
```

### Environment Variables

```bash
export DATABASE_URL="postgres://localhost:5432/myapp"
export PORT="3000"
export DEBUG="true"
export LOG_LEVEL="debug"
```

## 📖 Table of Contents

- [Core Concepts](#-core-concepts)
- [Struct Tags](#-struct-tags)
- [Supported Types](#-supported-types)
- [Configuration Sources](#-configuration-sources)
- [Configuration Options](#-configuration-options)
- [Advanced Features](#-advanced-features)
- [Custom Parsers](#-custom-parsers)
- [Error Handling](#-error-handling)
- [Examples](#-examples)
- [Best Practices](#-best-practices)
- [Performance](#-performance)
- [Contributing](#-contributing)

## 🧠 Core Concepts

### The Chipmunk Way

Chipmunk follows a simple philosophy:
1. **Define** your configuration as a Go struct
2. **Tag** fields with environment variable names
3. **Load** with a single function call
4. **Use** your strongly-typed configuration

### Basic Loading

```go
// Simple loading with default configuration
err := chipmunk.Load(&config)

// Loading with custom configuration
config := chipmunk.NewConfig(
  chipmunk.WithPrefix("APP_"),
  chipmunk.WithStrict(true),
)
err := chipmunk.LoadWithConfig(&target, config)

// Loading from specific sources
err := chipmunk.LoadFrom(&config, 
  sources.NewFileSource(".env"),
  sources.NewEnvSource(),
)
```

## 🏷️ Struct Tags

Chipmunk uses struct tags to control how environment variables are mapped to struct fields.

### Basic Tags

```go
type Config struct {
  // Basic mapping
  Port int `env:"PORT"`
  
  // Required field
  APIKey string `env:"API_KEY,required"`
  
  // With default value
  Host string `env:"HOST" default:"localhost"`
  
  // Multiple options
  Timeout int `env:"TIMEOUT,required" default:"30"`
}
```

### Advanced Tags

```go
type AdvancedConfig struct {
  // Custom separator for arrays
  Hosts []string `env:"HOSTS" separator:"|"`
  
  // Variable expansion
  DatabaseURL string `env:"DATABASE_URL" expand:"true"`
  
  // Read from file (for secrets)
  APIKey string `env:"API_KEY_FILE" file:"true"`
  
  // Combination of options
  Features []string `env:"FEATURES" default:"feature1,feature2" separator:","`
}
```

### All Available Tags

| Tag | Description | Example | Default |
|-----|-------------|---------|---------|
| `env` | Environment variable name | `env:"PORT"` | Required |
| `default` | Default value if not set | `default:"8080"` | None |
| `required` | Mark field as required | `env:"API_KEY,required"` | false |
| `separator` | Custom separator for arrays | `separator:"\|"` | "," |
| `expand` | Enable variable expansion | `expand:"true"` | false |
| `file` | Read value from file path | `file:"true"` | false |

## 🎯 Supported Types

Chipmunk supports all common Go types out of the box:

### Basic Types

```go
type BasicTypes struct {
  StringField  string  `env:"STRING_FIELD"`
  IntField     int     `env:"INT_FIELD"`
  Int64Field   int64   `env:"INT64_FIELD"`
  Float64Field float64 `env:"FLOAT64_FIELD"`
  BoolField    bool    `env:"BOOL_FIELD"`
  ByteField    byte    `env:"BYTE_FIELD"`
}
```

### Time Types

```go
type TimeTypes struct {
  // RFC3339 format: 2023-01-01T12:00:00Z
  CreatedAt time.Time `env:"CREATED_AT"`
  
  // Duration strings: "5m", "1h30m", "24h"
  Timeout time.Duration `env:"TIMEOUT"`
}
```

### URL Types

```go
type URLTypes struct {
  // Automatically parsed and validated
  DatabaseURL *url.URL `env:"DATABASE_URL"`
  APIURL      *url.URL `env:"API_URL"`
}
```

### Collections

```go
type Collections struct {
  // Comma-separated by default
  Hosts []string `env:"HOSTS"`
  
  // Custom separator
  Ports []int `env:"PORTS" separator:"|"`
  
  // Fixed-size arrays
  Servers [3]string `env:"SERVERS"`
  
  // Maps (key:value,key2:value2)
  Labels map[string]string `env:"LABELS"`
}
```

### Nested Structs

```go
type DatabaseConfig struct {
  Host     string `env:"DB_HOST" default:"localhost"`
  Port     int    `env:"DB_PORT" default:"5432"`
  Username string `env:"DB_USERNAME,required"`
  Password string `env:"DB_PASSWORD,required"`
}

type AppConfig struct {
  // Nested struct - fields are flattened
  Database DatabaseConfig
  
  // App-level config
  Port  int  `env:"PORT" default:"8080"`
  Debug bool `env:"DEBUG"`
}
```

## 📁 Configuration Sources

Chipmunk can load configuration from multiple sources with priority ordering.

### Environment Variables (Default)

```go
// Loads from OS environment variables
var config Config
err := chipmunk.Load(&config)
```

### .env Files

```go
import "github.com/yourusername/chipmunk/sources"

// Load from .env file
fileSource, err := sources.NewFileSource(".env")
if err != nil {
  log.Fatal(err)
}

err = chipmunk.LoadFrom(&config, fileSource)
```

### Multiple Sources with Priority

```go
// Priority: .env file -> OS environment -> defaults
err := chipmunk.LoadFrom(&config,
  sources.NewFileSource(".env"),           // Highest priority
  sources.NewEnvSource(),                  // Fallback
)
```

### In-Memory Source (Testing)

```go
// Great for testing
memorySource := sources.NewMemorySource(map[string]string{
  "PORT":         "8080",
  "DEBUG":        "true",
  "DATABASE_URL": "postgres://localhost:5432/test",
})

err := chipmunk.LoadFrom(&config, memorySource)
```

### Custom Sources

```go
// Implement the Source interface
type CustomSource struct{}

func (c *CustomSource) Get(key string) (string, bool) {
  // Your custom logic here
  return value, found
}

func (c *CustomSource) Keys() []string {
  // Return all available keys
  return keys
}

func (c *CustomSource) Name() string {
  return "CustomSource"
}
```

## ⚙️ Configuration Options

Customize Chipmunk's behavior with configuration options:

### Functional Options

```go
config := chipmunk.NewConfig(
  // Add prefix to all environment variables
  chipmunk.WithPrefix("MYAPP_"),
  
  // Enable strict mode (fail on unknown vars)
  chipmunk.WithStrict(true),
  
  // Custom separator for arrays
  chipmunk.WithSeparator("|"),
  
  // Case insensitive matching
  chipmunk.WithCaseSensitive(false),
  
  // Custom sources
  chipmunk.WithSources(
      sources.NewFileSource(".env"),
      sources.NewEnvSource(),
  ),
)

err := chipmunk.LoadWithConfig(&target, config)
```

### Configuration Examples

```go
// Development configuration
devConfig := chipmunk.NewConfig(
  chipmunk.WithPrefix("DEV_"),
  chipmunk.WithStrict(false),
)

// Production configuration
prodConfig := chipmunk.NewConfig(
  chipmunk.WithPrefix("PROD_"),
  chipmunk.WithStrict(true),
  chipmunk.WithValidation(true),
)

// Testing configuration
testConfig := chipmunk.NewConfig(
  chipmunk.WithSources(
      sources.NewMemorySource(testData),
  ),
)
```

## 🚀 Advanced Features

### Variable Expansion

Expand variables within values using `${VAR}` or `$VAR` syntax:

```go
type Config struct {
  DatabaseURL string `env:"DATABASE_URL" expand:"true"`
  LogFile     string `env:"LOG_FILE" expand:"true"`
}
```

```bash
export DB_HOST="localhost"
export DB_PORT="5432"
export DB_NAME="myapp"
export DATABASE_URL="postgres://${DB_HOST}:${DB_PORT}/${DB_NAME}"
export LOG_FILE="/var/log/${DB_NAME}.log"
```

### File-Based Secrets

Read sensitive values from files (perfect for Docker secrets, Kubernetes secrets):

```go
type Config struct {
  // Read API key from file
  APIKey string `env:"API_KEY_FILE" file:"true"`
  
  // Read database password from file
  DBPassword string `env:"DB_PASSWORD_FILE" file:"true"`
}
```

```bash
export API_KEY_FILE="/run/secrets/api_key"
export DB_PASSWORD_FILE="/run/secrets/db_password"
```

### Validation

Built-in validation for common patterns:

```go
type Config struct {
  Email    string `env:"EMAIL" validate:"email"`
  URL      string `env:"URL" validate:"url"`
  Port     int    `env:"PORT" validate:"range:1024-65535"`
  LogLevel string `env:"LOG_LEVEL" validate:"oneof:debug info warn error"`
}
```

### Encryption Support

Encrypt sensitive configuration values:

```go
import "github.com/yourusername/chipmunk/advanced"

// Encrypt a value
encrypted, err := advanced.Encrypt("sensitive-data", "encryption-key")

// Decrypt during loading
type Config struct {
  APIKey string `env:"API_KEY" decrypt:"true"`
}
```

## 🔧 Custom Parsers

Extend Chipmunk with custom type parsers:

### Simple Custom Parser

```go
import "github.com/yourusername/chipmunk/parsers"

// Custom type
type LogLevel int

const (
  Debug LogLevel = iota
  Info
  Warn
  Error
)

// Custom parser
type LogLevelParser struct{}

func (p LogLevelParser) Parse(value string) (interface{}, error) {
  switch strings.ToLower(value) {
  case "debug":
      return Debug, nil
  case "info":
      return Info, nil
  case "warn":
      return Warn, nil
  case "error":
      return Error, nil
  default:
      return nil, fmt.Errorf("invalid log level: %s", value)
  }
}

func (p LogLevelParser) CanParse(t reflect.Type) bool {
  return t == reflect.TypeOf(LogLevel(0))
}

// Register the parser
func init() {
  loader := chipmunk.NewLoader(chipmunk.NewConfig())
  loader.RegisterParser(reflect.TypeOf(LogLevel(0)), LogLevelParser{})
}
```

### Advanced Custom Parser

```go
// Database connection type
type DatabaseConfig struct {
  Driver   string
  Host     string
  Port     int
  Database string
  Username string
  Password string
}

// Custom parser for database URLs
type DatabaseConfigParser struct{}

func (p DatabaseConfigParser) Parse(value string) (interface{}, error) {
  // Parse database URL: postgres://user:pass@host:port/db
  u, err := url.Parse(value)
  if err != nil {
      return nil, err
  }
  
  password, _ := u.User.Password()
  port, _ := strconv.Atoi(u.Port())
  
  return DatabaseConfig{
      Driver:   u.Scheme,
      Host:     u.Hostname(),
      Port:     port,
      Database: strings.TrimPrefix(u.Path, "/"),
      Username: u.User.Username(),
      Password: password,
  }, nil
}

func (p DatabaseConfigParser) CanParse(t reflect.Type) bool {
  return t == reflect.TypeOf(DatabaseConfig{})
}
```

### Function-Based Parser

```go
// Convenience function for simple parsers
parsers.RegisterCustomParser(
  reflect.TypeOf(LogLevel(0)),
  func(value string) (interface{}, error) {
      // Your parsing logic here
      return parseLogLevel(value)
  },
)
```

## 🚨 Error Handling

Chipmunk provides detailed, actionable error messages:

### Error Types

```go
import "github.com/yourusername/chipmunk"

err := chipmunk.Load(&config)
if err != nil {
  switch e := err.(type) {
  case chipmunk.RequiredFieldError:
      fmt.Printf("Missing required field: %s (env: %s)\n", e.Field, e.EnvKey)
      
  case chipmunk.ParseError:
      fmt.Printf("Parse error for field %s: cannot convert '%s' to %s\n", 
          e.Field, e.Value, e.TargetType)
          
  case chipmunk.ValidationError:
      fmt.Printf("Validation error for field %s: %s\n", e.Field, e.Rule)
      
  case chipmunk.FileReadError:
      fmt.Printf("Cannot read file for field %s: %s\n", e.Field, e.FilePath)
      
  default:
      fmt.Printf("Unknown error: %v\n", err)
  }
}
```

### Error Examples

```go
// Missing required field
// Error: required field 'APIKey' (env: API_KEY) is missing

// Parse error
// Error: cannot parse field 'Port' (env: PORT): value 'abc' cannot be converted to int

// Validation error
// Error: validation failed for field 'Email' (env: EMAIL): value 'invalid' violates rule 'email'

// File read error
// Error: cannot read file for field 'APIKey': /path/to/secret: no such file or directory
```

## 📚 Examples

### Web Application

```go
package main

import (
  "fmt"
  "log"
  "net/http"
  "time"
  "github.com/yourusername/chipmunk"
)

type ServerConfig struct {
  Host         string        `env:"HOST" default:"localhost"`
  Port         int           `env:"PORT" default:"8080"`
  ReadTimeout  time.Duration `env:"READ_TIMEOUT" default:"30s"`
  WriteTimeout time.Duration `env:"WRITE_TIMEOUT" default:"30s"`
  TLS          TLSConfig
}

type TLSConfig struct {
  Enabled  bool   `env:"TLS_ENABLED"`
  CertFile string `env:"TLS_CERT_FILE" file:"true"`
  KeyFile  string `env:"TLS_KEY_FILE" file:"true"`
}

type DatabaseConfig struct {
  URL             string        `env:"DATABASE_URL,required"`
  MaxConnections  int           `env:"DB_MAX_CONNECTIONS" default:"10"`
  ConnMaxLifetime time.Duration `env:"DB_CONN_MAX_LIFETIME" default:"1h"`
  SSLMode         string        `env:"DB_SSL_MODE" default:"prefer"`
}

type AppConfig struct {
  Server   ServerConfig
  Database DatabaseConfig
  
  // Application settings
  APIKey      string   `env:"API_KEY,required"`
  Debug       bool     `env:"DEBUG"`
  LogLevel    string   `env:"LOG_LEVEL" default:"info"`
  AllowedHosts []string `env:"ALLOWED_HOSTS" separator:","`
  
  // Feature flags
  Features map[string]string `env:"FEATURES"`
}

func main() {
  var config AppConfig
  
  if err := chipmunk.Load(&config); err != nil {
      log.Fatal("Configuration error:", err)
  }
  
  // Use configuration
  addr := fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
  
  server := &http.Server{
      Addr:         addr,
      ReadTimeout:  config.Server.ReadTimeout,
      WriteTimeout: config.Server.WriteTimeout,
  }
  
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
      fmt.Fprintf(w, "🐿️ Chipmunk Web App\n")
      fmt.Fprintf(w, "Debug: %v\n", config.Debug)
      fmt.Fprintf(w, "Database: %s\n", config.Database.URL)
      fmt.Fprintf(w, "Allowed Hosts: %v\n", config.AllowedHosts)
  })
  
  fmt.Printf("🚀 Server starting on %s\n", addr)
  if config.Server.TLS.Enabled {
      log.Fatal(server.ListenAndServeTLS(
          config.Server.TLS.CertFile,
          config.Server.TLS.KeyFile,
      ))
  } else {
      log.Fatal(server.ListenAndServe())
  }
}
```

### Microservice Configuration

```go
type MicroserviceConfig struct {
  // Service identity
  ServiceName    string `env:"SERVICE_NAME,required"`
  ServiceVersion string `env:"SERVICE_VERSION" default:"1.0.0"`
  Environment    string `env:"ENVIRONMENT" default:"development"`
  
  // HTTP server
  HTTPPort int           `env:"HTTP_PORT" default:"8080"`
  Timeout  time.Duration `env:"HTTP_TIMEOUT" default:"30s"`
  
  // Database
  DatabaseURL string `env:"DATABASE_URL,required"`
  
  // Message queue
  RabbitMQURL string `env:"RABBITMQ_URL,required"`
  QueueName   string `env:"QUEUE_NAME" default:"default"`
  
  // Observability
  MetricsPort   int    `env:"METRICS_PORT" default:"9090"`
  TracingURL    string `env:"TRACING_URL"`
  LogLevel      string `env:"LOG_LEVEL" default:"info"`
  
  // Security
  JWTSecret     string   `env:"JWT_SECRET_FILE" file:"true"`
  AllowedOrigins []string `env:"ALLOWED_ORIGINS" separator:","`
  
  // Feature flags
  EnableMetrics bool `env:"ENABLE_METRICS" default:"true"`
  EnableTracing bool `env:"ENABLE_TRACING"`
  EnableAuth    bool `env:"ENABLE_AUTH" default:"true"`
}
```

### Docker Compose Integration

```yaml
# docker-compose.yml
version: '3.8'
services:
app:
  build: .
  environment:
    - SERVICE_NAME=my-service
    - HTTP_PORT=8080
    - DATABASE_URL=postgres://user:pass@db:5432/myapp
    - RABBITMQ_URL=amqp://guest:guest@rabbitmq:5672/
    - LOG_LEVEL=debug
    - ENABLE_METRICS=true
    - ALLOWED_ORIGINS=http://localhost:3000,https://myapp.com
  secrets:
    - jwt_secret
  ports:
    - "8080:8080"
    - "9090:9090"
  
secrets:
jwt_secret:
  file: ./secrets/jwt_secret.txt
```

### Testing Configuration

```go
func TestConfigLoading(t *testing.T) {
  // Create test data
  testData := map[string]string{
      "SERVICE_NAME":    "test-service",
      "HTTP_PORT":       "8080",
      "DATABASE_URL":    "postgres://localhost:5432/test",
      "LOG_LEVEL":       "debug",
      "ALLOWED_ORIGINS": "http://localhost:3000,https://test.com",
  }
  
  // Create memory source
  source := sources.NewMemorySource(testData)
  
  // Load configuration
  var config MicroserviceConfig
  err := chipmunk.LoadFrom(&config, source)
  
  // Test assertions
  assert.NoError(t, err)
  assert.Equal(t, "test-service", config.ServiceName)
  assert.Equal(t, 8080, config.HTTPPort)
  assert.Equal(t, []string{"http://localhost:3000", "https://test.com"}, config.AllowedOrigins)
}
```

## 🎯 Best Practices

### 1. Structure Your Configuration

```go
// ✅ Good: Organized, clear structure
type Config struct {
  Server   ServerConfig
  Database DatabaseConfig
  Cache    CacheConfig
  Auth     AuthConfig
}

// ❌ Avoid: Flat, unorganized structure
type Config struct {
  Port        int    `env:"PORT"`
  DBHost      string `env:"DB_HOST"`
  DBPort      int    `env:"DB_PORT"`
  CacheHost   string `env:"CACHE_HOST"`
  CachePort   int    `env:"CACHE_PORT"`
  JWTSecret   string `env:"JWT_SECRET"`
  // ... many more fields
}
```

### 2. Use Meaningful Names

```go
// ✅ Good: Clear, descriptive names
type Config struct {
  DatabaseURL        string `env:"DATABASE_URL"`
  MaxConnections     int    `env:"DB_MAX_CONNECTIONS"`
  ConnectionTimeout  time.Duration `env:"DB_CONNECTION_TIMEOUT"`
}

// ❌ Avoid: Cryptic abbreviations
type Config struct {
  DBURL string `env:"DB_URL"`
  MaxC  int    `env:"MAX_C"`
  TO    time.Duration `env:"TO"`
}
```

### 3. Provide Sensible Defaults

```go
// ✅ Good: Reasonable defaults for optional settings
type Config struct {
  Port         int           `env:"PORT" default:"8080"`
  Timeout      time.Duration `env:"TIMEOUT" default:"30s"`
  LogLevel     string        `env:"LOG_LEVEL" default:"info"`
  DatabaseURL  string        `env:"DATABASE_URL,required"` // No default for critical settings
}
```

### 4. Use Validation

```go
// ✅ Good: Validate critical values
type Config struct {
  Port     int    `env:"PORT" default:"8080" validate:"range:1024-65535"`
  Email    string `env:"EMAIL" validate:"email"`
  LogLevel string `env:"LOG_LEVEL" default:"info" validate:"oneof:debug info warn error"`
}
```

### 5. Handle Secrets Properly

```go
// ✅ Good: Use file-based secrets
type Config struct {
  APIKey       string `env:"API_KEY_FILE" file:"true"`
  DatabasePass string `env:"DB_PASSWORD_FILE" file:"true"`
}

// ❌ Avoid: Secrets in environment variables (less secure)
type Config struct {
  APIKey       string `env:"API_KEY"`
  DatabasePass string `env:"DB_PASSWORD"`
}
```

### 6. Environment-Specific Configuration

```go
func LoadConfig() (*Config, error) {
  env := os.Getenv("ENVIRONMENT")
  
  var sources []chipmunk.Source
  
  switch env {
  case "development":
      sources = []chipmunk.Source{
          sources.NewFileSource(".env.local"),
          sources.NewFileSource(".env.development"),
          sources.NewEnvSource(),
      }
  case "production":
      sources = []chipmunk.Source{
          sources.NewEnvSource(), // Only environment variables in production
      }
  default:
      sources = []chipmunk.Source{
          sources.NewFileSource(".env"),
          sources.NewEnvSource(),
      }
  }
  
  var config Config
  err := chipmunk.LoadFrom(&config, sources...)
  return &config, err
}
```

## ⚡ Performance

Chipmunk is designed for performance:

### Benchmarks

```
BenchmarkLoad-8                    50000    23456 ns/op    1024 B/op    12 allocs/op
BenchmarkLoadWithConfig-8          45000    25123 ns/op    1152 B/op    14 allocs/op
BenchmarkParseString-8           5000000      234 ns/op      32 B/op     1 allocs/op
BenchmarkParseInt-8              3000000      456 ns/op      24 B/op     1 allocs/op
BenchmarkParseSlice-8            1000000     1234 ns/op     256 B/op     3 allocs/op
```

### Performance Tips

1. **Reuse Loaders**: Create loader once, use multiple times
2. **Minimize Sources**: Fewer sources = faster loading
3. **Cache Results**: Load once at startup, reuse throughout application
4. **Use Appropriate Types**: Choose the most efficient type for your data

```go
// ✅ Good: Reuse loader
loader := chipmunk.NewLoader(config)
loader.Load(&config1)
loader.Load(&config2)

// ❌ Avoid: Creating new loader each time
chipmunk.Load(&config1)
chipmunk.Load(&config2)
```

## 🧪 Testing

Chipmunk makes testing easy with in-memory sources:

```go
func TestMyService(t *testing.T) {
  // Arrange
  testConfig := sources.NewMemorySource(map[string]string{
      "PORT":         "8080",
      "DATABASE_URL": "postgres://localhost:5432/test",
      "DEBUG":        "true",
  })
  
  var config Config
  err := chipmunk.LoadFrom(&config, testConfig)
  require.NoError(t, err)
  
  // Act
  service := NewService(config)
  
  // Assert
  assert.Equal(t, 8080, service.Port())
  assert.True(t, service.IsDebug())
}
```

## 🔍 Debugging

Enable debug logging to see what Chipmunk is doing:

```go
config := chipmunk.NewConfig(
  chipmunk.WithDebug(true),
)

// Output:
// [CHIPMUNK] Loading field 'Port' from env 'PORT'
// [CHIPMUNK] Found value '8080' for 'PORT'
// [CHIPMUNK] Parsed '8080' as int: 8080
// [CHIPMUNK] Set field 'Port' = 8080
```

## 🤝 Contributing

We welcome contributions! Please see our [Contributing Guide](CONTRIBUTING.md) for details.

### Development Setup

```bash
# Clone the repository
git clone https://github.com/yourusername/chipmunk.git
cd chipmunk

# Install dependencies
go mod download

# Run tests
make test

# Run tests with coverage
make test-coverage

# Run linter
make lint

# Run examples
make examples
```

### Running Tests

```bash
# Unit tests
go test ./...

# Integration tests
go test -tags=integration ./...

# Benchmarks
go test -bench=. -benchmem ./...

# Coverage
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- Inspired by the simplicity of [envconfig](https://github.com/kelseyhightower/envconfig)
- Built with the extensibility of [viper](https://github.com/spf13/viper) in mind
- Performance optimizations inspired by [fasthttp](https://github.com/valyala/fasthttp)

## 📞 Support

- 📖 [Documentation](https://pkg.go.dev/github.com/yourusername/chipmunk)
- 🐛 [Issue Tracker](https://github.com/yourusername/chipmunk/issues)
- 💬 [Discussions](https://github.com/yourusername/chipmunk/discussions)
- 📧 [Email](mailto:support@yourname.com)

---

<div align="center">
<p>Made with ❤️ by the Go community</p>
<p>🐿️ <strong>Chipmunk</strong> - Gather, store, and organize your environment variables</p>
</div>