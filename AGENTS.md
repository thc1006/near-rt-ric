# Near-RT RIC Platform Configuration

## Project Overview

**Near-RT RIC (Near Real-Time RAN Intelligent Controller)** - A comprehensive O-RAN compliant platform implementing intelligent RAN optimization through network slicing, federated learning, and multi-functional xApps. This project combines Go backend services (40.3%) with Angular frontend applications (30%) to deliver production-ready O-RAN solutions.

### Architecture Context
- **Near-RT RIC Platform**: Handles near-real-time RAN control and optimization (10ms-1s latency requirements)
- **xApp Ecosystem**: Hosts applications providing specific RAN management functions via E2 interface
- **Federated Learning Integration**: Implements privacy-preserving ML models for intelligent RRM
- **Network Slicing**: Advanced capabilities with dynamic slice creation and SLA monitoring
- **Multi-vendor Support**: Compatible with different RAN vendor implementations

## Technology Stack & Language Preferences

### Backend (Go - 40.3%)
- **Primary Language**: Go 1.21+
- **Framework**: Standard library with net/http, gorilla/mux for routing
- **Concurrency**: Use Go channels and goroutines for high-throughput scenarios
- **Error Handling**: Always implement proper error handling with context.Context
- **Logging**: Use structured logging (logrus or zap) with appropriate log levels
- **Testing**: Use testify framework with gomock for mocking

### Frontend (Angular/TypeScript - 30%)
- **Framework**: Angular 17+ with standalone components
- **Language**: TypeScript with strict mode enabled
- **UI Library**: Angular Material + CDK for consistent components
- **State Management**: Use NgRx for complex state, signals for simple state
- **Testing**: Jest for unit tests, Cypress for e2e tests
- **Styling**: SCSS with consistent design system

### Infrastructure
- **Containerization**: Docker with multi-stage builds
- **Orchestration**: Kubernetes with proper RBAC and network policies
- **Message Routing**: RMR (RIC Message Router) for xApp communication
- **Database**: Support for time-series data (InfluxDB) and relational (PostgreSQL)

## Code Style & Standards

### Go Code Standards
```
// Use structured logging with context
log.WithFields(logrus.Fields{
    "xapp_id": xappID,
    "subscription_id": subID,
}).Info("Processing E2 subscription request")

// Error handling with context
if err := processE2Message(ctx, msg); err != nil {
    return fmt.Errorf("failed to process E2 message: %w", err)
}

// Interface definitions for xApp communication
type XAppManager interface {
    RegisterXApp(ctx context.Context, xapp *XAppConfig) error
    DeregisterXApp(ctx context.Context, xappID string) error
    RouteMessage(ctx context.Context, msg *E2Message) error
}
```

### TypeScript/Angular Standards
```
// Use strict typing with proper interfaces
interface E2NodeStatus {
  nodeId: string;
  status: 'connected' | 'disconnected' | 'error';
  lastHeartbeat: Date;
  cellInfo: CellConfiguration[];
}

// Component with proper lifecycle management
@Component({
  selector: 'app-e2-node-monitor',
  standalone: true,
  imports: [CommonModule, MatTableModule]
})
export class E2NodeMonitorComponent implements OnInit, OnDestroy {
  // Use signals for reactive state
  nodeStatus = signal([]);
}
```

### Naming Conventions
- **Go**: CamelCase for exported functions, camelCase for internal
- **TypeScript**: camelCase for variables and functions, PascalCase for types/interfaces
- **Files**: kebab-case for component files, snake_case for Go packages
- **Constants**: UPPER_SNAKE_CASE for both languages
- **O-RAN Terminology**: Use standard O-RAN terms (xApp, E2Node, RIC, etc.)

## Build & Development Commands

### Local Development Setup
```
# Full development environment startup
make dev-setup          # Initialize development environment
make install-deps       # Install all dependencies (Go modules, npm packages)
make generate-proto     # Generate protobuf files for E2 interface
make build-dev          # Build development versions of all components

# Component-specific development
make dashboard-dev      # Start dashboard development server (localhost:4200)
make xapp-dashboard-dev # Start xApp dashboard development server (localhost:4201)
make ric-server-dev     # Start RIC backend server (localhost:8080)
```

### Build Commands
```
# Production builds
make build              # Build all components for production
make build-go           # Build Go backend services only
make build-angular      # Build Angular frontends only
make build-docker       # Build Docker images for all services

# Clean builds
make clean              # Clean all build artifacts
make clean-go           # Clean Go build cache
make clean-node         # Clean node_modules and package locks
```

### Testing Commands
```
# Comprehensive testing
make test               # Run all tests (unit, integration, e2e)
make test-go            # Run Go tests with coverage
make test-angular       # Run Angular unit tests (Jest)
make test-e2e           # Run end-to-end tests (Cypress)
make test-integration   # Run integration tests with mock RAN

# Coverage requirements
make coverage           # Generate test coverage reports (minimum 85%)
make coverage-report    # Generate HTML coverage reports
```

### Quality Assurance
```
# Code quality checks
make lint               # Run all linters
make lint-go            # Run Go linters (golangci-lint)
make lint-ts            # Run TypeScript/Angular linters (ESLint)
make format             # Auto-format all code
make security-scan      # Run security vulnerability scans
```

## Architecture & Design Patterns

### O-RAN Component Architecture
```
Near-RT RIC Platform
├── A1 Interface (Policy Management)
├── E2 Interface (RAN Communication)
├── xApp Framework (Application Hosting)
├── RMR (Message Routing)
├── SDL (Shared Data Layer)
└── Dashboard Interface (Management UI)
```

### Code Organization Patterns
- **Go Services**: Follow clean architecture with handlers, services, repositories
- **Angular Apps**: Feature-based modules with lazy loading
- **xApp Development**: Standard xApp template with E2 subscription management
- **API Design**: RESTful APIs with OpenAPI 3.0 specifications
- **Configuration**: Environment-based config with validation

### Key Interfaces to Implement
```
// E2 Interface for RAN communication
type E2Interface interface {
    Subscribe(ctx context.Context, req *E2SubscriptionRequest) error
    Unsubscribe(ctx context.Context, subID string) error
    SendControlMessage(ctx context.Context, msg *E2ControlMessage) error
}

// A1 Interface for policy management
type A1Interface interface {
    CreatePolicy(ctx context.Context, policy *A1Policy) error
    UpdatePolicy(ctx context.Context, policyID string, policy *A1Policy) error
    DeletePolicy(ctx context.Context, policyID string) error
}
```

## Testing Strategy & Requirements

### Test Coverage Requirements
- **Go Backend**: Minimum 90% coverage for critical RIC functions
- **Angular Frontend**: Minimum 85% coverage for components and services
- **Integration Tests**: Cover all O-RAN interface protocols
- **E2E Tests**: Simulate complete xApp lifecycle scenarios

### Test Categories
```
# Unit Tests
go test ./... -v -race -coverprofile=coverage.out
npm run test -- --coverage --watchAll=false

# Integration Tests (with mock RAN simulator)
make test-e2-integration    # Test E2 interface with simulated nodes
make test-a1-integration    # Test A1 interface with policy scenarios
make test-xapp-lifecycle    # Test complete xApp deployment/management

# Performance Tests
make test-latency          # Verify 10ms-1s control loop requirements
make test-throughput       # Test concurrent xApp message handling
make test-stress           # Stress test with multiple E2 nodes
```

### Mock Configurations
- Use RAN simulators for E2 interface testing
- Mock federated learning endpoints for ML testing
- Simulate network slice scenarios for slicing tests
- Test with multi-vendor E2 service models

## Security & Compliance

### O-RAN Security Requirements
- **Authentication**: Implement O-RAN compliant authentication mechanisms
- **TLS Configuration**: Use TLS 1.3 for all external communications
- **Certificate Management**: Proper cert rotation for E2 and A1 interfaces
- **Data Protection**: Encrypt sensitive RAN data and UE information
- **Access Control**: Role-based access for different operator personas

### Security Scanning & Validation
```
make security-full-scan    # Complete security vulnerability assessment
make compliance-check      # Verify O-RAN Alliance compliance
make secret-scan          # Check for exposed credentials/secrets
make container-scan       # Scan Docker images for vulnerabilities
```

## Deployment & Operations

### Container & Kubernetes Configuration
```
# Local deployment
make deploy-local         # Deploy to local Kubernetes cluster
make deploy-minikube     # Deploy to minikube for development

# Production deployment
make deploy-prod         # Deploy production-ready configuration
make deploy-ha           # Deploy high-availability configuration
make deploy-monitoring   # Deploy Prometheus/Grafana monitoring stack
```

### Environment Management
- **Development**: Single-node deployment with mock RAN
- **Testing**: Multi-node deployment with RAN simulators
- **Production**: HA deployment with real RAN integration
- **Configuration**: Use Helm charts for environment-specific configs

### Monitoring & Observability
```
# Monitoring setup
make setup-monitoring     # Deploy Prometheus, Grafana, Jaeger
make setup-logging       # Deploy ELK stack for centralized logging
make setup-alerting      # Configure alerting rules for O-RAN KPIs
```

## Federated Learning & AI/ML Integration

### ML Model Management
- Use TensorFlow Serving or TorchServe for model deployment
- Implement model versioning and A/B testing capabilities
- Support distributed training across multiple xApps
- Ensure privacy-preserving aggregation algorithms

### Data Pipeline Requirements
```
make ml-pipeline-setup   # Setup ML training and inference pipelines
make fl-test            # Test federated learning aggregation
make model-validate     # Validate ML model accuracy and performance
```

## Documentation & PR Standards

### Required Documentation Updates
When modifying code, ALWAYS update:
- API documentation (OpenAPI specs in `/docs/api/`)
- Architecture diagrams (Mermaid/PlantUML in `/docs/architecture/`)
- O-RAN compliance matrices (in `/docs/compliance/`)
- Security assessment reports (in `/docs/security/`)

### Pull Request Requirements
```
# PR Template (required sections)
## Summary
Brief description of changes and motivation

## O-RAN Compliance Impact
- [ ] No impact on O-RAN compliance
- [ ] Requires compliance review
- [ ] Updates compliance documentation

## Testing Completed
- [ ] Unit tests pass (90%+ coverage)
- [ ] Integration tests with RAN simulator
- [ ] Performance tests within SLA
- [ ] Security scan completed

## Deployment Impact
- [ ] No deployment changes required
- [ ] Requires configuration updates
- [ ] Requires database migration
- [ ] Affects multiple components
```

### Commit Message Format
```
<type>(<scope>): <subject>

[optional body]

[optional footer(s)]

Examples:
feat(e2): add support for E2SM-RC service model
fix(dashboard): resolve xApp registration timeout issue
docs(api): update E2 interface OpenAPI specification
```

## Performance & Scalability Requirements

### Latency Requirements
- **E2 Control Messages**: < 10ms processing time
- **A1 Policy Updates**: < 100ms propagation time
- **Dashboard Response**: < 200ms for UI interactions
- **xApp Registration**: < 1s complete lifecycle

### Throughput Requirements
- Support 1000+ concurrent E2 subscriptions
- Handle 10,000+ E2 indications per second
- Manage 100+ simultaneous xApp instances
- Process federated learning updates from 50+ participants

### Resource Management
```
# Performance monitoring
make perf-monitor        # Monitor system performance metrics
make latency-test       # Measure control loop latency
make throughput-test    # Measure message processing throughput
make resource-usage     # Monitor CPU/memory usage patterns
```

## Environment Variables & Configuration

### Required Environment Variables
```
# RIC Platform Configuration
export RIC_PLATFORM_NAMESPACE="ricplt"
export RIC_XAPP_NAMESPACE="ricxapp"
export RIC_INFRA_NAMESPACE="ricinfra"

# E2 Interface Configuration
export E2_TERM_HOST="e2term-service"
export E2_TERM_PORT="38000"
export E2_SUB_HOST="subscription-service"
export E2_SUB_PORT="8088"

# A1 Interface Configuration
export A1_MEDIATOR_HOST="a1mediator-service"
export A1_MEDIATOR_PORT="10000"

# Database Configuration
export RIC_DB_HOST="postgresql-service"
export RIC_DB_PORT="5432"
export RIC_DB_NAME="ricplt"

# Security Configuration
export TLS_CERT_PATH="/opt/ric/certs"
export JWT_SECRET_KEY="your-jwt-secret"
export ENCRYPTION_KEY="your-encryption-key"
```

### Configuration Validation
```
make validate-config    # Validate all configuration files
make check-env         # Check required environment variables
make test-connectivity # Test connectivity to external services
```

## Troubleshooting & Debugging

### Common Issues & Solutions
```
# E2 Interface Issues
make debug-e2          # Debug E2 connection and subscription issues
make test-e2-sim       # Test with E2 RAN simulator

# xApp Management Issues
make debug-xapp        # Debug xApp registration and lifecycle
make logs-xapp         # Collect xApp logs for analysis

# Performance Issues
make profile-cpu       # CPU profiling for Go services
make profile-memory    # Memory profiling for Go services
make trace-requests    # Distributed tracing for request flows
```

### Log Collection & Analysis
```
make collect-logs      # Collect logs from all components
make analyze-logs      # Analyze logs for errors and patterns
make export-metrics    # Export metrics for external analysis
```

## Final Notes for AI Agents

### Critical Success Factors
1. **O-RAN Compliance**: Always verify changes maintain O-RAN Alliance specification compliance
2. **Latency Requirements**: Ensure all changes respect near-real-time performance requirements
3. **Security First**: Never compromise on security, especially for RAN data handling
4. **Multi-vendor Support**: Test changes across different vendor implementations
5. **Production Readiness**: All code must be production-ready with proper error handling

### When Modifying This Project
- Read and understand O-RAN specifications before making architectural changes
- Test with RAN simulators before proposing changes to E2/A1 interfaces
- Validate federated learning implementations with privacy-preserving requirements
- Ensure all changes support multi-vendor interoperability
- Document any new O-RAN features or capabilities added

### Communication Protocols
- Follow O-RAN defined message formats for E2 and A1 interfaces
- Implement proper ASN.1 encoding/decoding for E2 messages
- Use standard RMR messaging patterns for xApp communication
- Maintain backward compatibility with existing xApp implementations

This configuration ensures that AI agents working on this Near-RT RIC platform understand the complex O-RAN ecosystem requirements and can generate code that meets telecommunications industry standards while maintaining the intelligent features of federated learning and network slicing.
