# Learning Module gRPC Service

A high-performance gRPC service that acts as a bridge between Go clients and a Django REST API for managing learning modules. 

## Tech Stack

- Go 1.21+
- Protocol Buffers (protobuf)
- gRPC
- HTTP client for Django API communication

## Project Structure

```
go-service/
├── cmd/
│   └── server/
│       └── main.go         # Server entry point
├── client/
│   └── main.go            # Example client implementation
├── pb/
│   ├── module.pb.go       # Generated protobuf code
│   └── module_grpc.pb.go  # Generated gRPC code
├── proto/
│   └── module.proto       # Protocol buffer definitions
├── server/
│   └── server.go          # Server implementation
├── go.mod                 # Go module file
└── go.sum                 # Go module checksums
```

## Prerequisites

- Go 1.21 or later
- Protocol Buffers compiler (protoc)
- Go plugins for protoc:
  - protoc-gen-go
  - protoc-gen-go-grpc
- Running Django REST API (http://localhost:8000)

## Installation

1. Clone the repository:
   ```bash
   git clone <repository-url>
   cd go-service
   ```

2. Install Protocol Buffers compiler:
   ```bash
   # Windows (using chocolatey)
   choco install protoc

   # macOS
   brew install protobuf

   # Linux
   sudo apt install -y protobuf-compiler
   ```

3. Install Go plugins for protoc:
   ```bash
   go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
   ```

4. Install Go dependencies:
   ```bash
   go mod tidy
   ```

5. Generate protobuf code:
   ```bash
   protoc --go_out=. --go_opt=paths=source_relative \
          --go-grpc_out=. --go-grpc_opt=paths=source_relative \
          proto/module.proto
   ```

## Running the Service

1. Ensure the Django REST API is running:
   ```bash
   # In the Django project directory
   python manage.py runserver
   ```

2. Start the gRPC server:
   ```bash
   # In the go-service directory
   go run cmd/server/main.go
   ```

3. Run the example client:
   ```bash
   # In a new terminal
   go run client/main.go
   ```


### Message Types

- `Empty`: Empty request for ListModules
- `ModuleRequest`: Request with module ID for GetModule
- `Module`: Module data structure
- `ModuleList`: List of modules

### Module Structure

```protobuf
message Module {
    int32 id = 1;
    string subject = 2;
    string grade = 3;
    string chapter = 4;
    string content_json = 5;
}
```

## Development

### Code Generation

After modifying `module.proto`:
```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/module.proto
```

### Testing

Run the example client to test the service:
```bash

# you can read all modules by this
go run client/main.go

# if you want specific module by ID (e.g., ID 1)
go run client/main.go 1
```