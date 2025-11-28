# gRPC Go Tutorial Project

A comprehensive gRPC tutorial project demonstrating all four types of RPC communication patterns in Go.

## 📚 What is gRPC?

gRPC (Google Remote Procedure Call) is a modern, high-performance RPC framework that enables client and server applications to communicate transparently. It uses Protocol Buffers (protobuf) for serialization and supports multiple programming languages.

## 🎯 Project Overview

This project implements a **GreetService** with four different RPC patterns:

1. **Unary RPC** - Simple request/response (1 request → 1 response)
2. **Server Streaming** - Client sends one request, server sends multiple responses
3. **Client Streaming** - Client sends multiple requests, server sends one response
4. **Bidirectional Streaming** - Both client and server send multiple messages

## 📁 Project Structure

```
grpc-go/
├── proto/                    # Protocol buffer definitions
│   ├── greet.proto          # Service and message definitions
│   ├── greet.pb.go          # Generated message types
│   └── greet_grpc.pb.go     # Generated service code
├── server/                  # gRPC server implementation
│   ├── main.go             # Server setup and registration
│   ├── unary.go            # Unary RPC implementation
│   ├── server_stream.go    # Server streaming implementation
│   ├── client_stream.go    # Client streaming implementation
│   └── bi_stream.go        # Bidirectional streaming implementation
├── client/                  # gRPC client implementation
│   ├── main.go             # Client setup
│   ├── unary.go            # Unary RPC client call
│   ├── server_stream.go    # Server streaming client call
│   ├── client_stream.go    # Client streaming client call
│   └── bi_stream.go        # Bidirectional streaming client call
├── go.mod                   # Go module dependencies
└── README.md               # This file
```

## 🚀 Prerequisites

- **Go** 1.24+ installed ([Download](https://go.dev/dl/))
- **Protocol Buffer Compiler (protoc)** installed
  - macOS: `brew install protobuf`
  - Linux: `sudo apt-get install protobuf-compiler`
  - Windows: Download from [protobuf releases](https://github.com/protocolbuffers/protobuf/releases)
- **Go plugins for protoc**:
  ```bash
  go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
  go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
  ```

## ⚙️ Setup

1. **Clone the repository**:
   ```bash
   git clone https://github.com/harsh2971/grpc-go.git
   cd grpc-go
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Generate Go code from proto files**:
   ```bash
   protoc --go_out=. --go_opt=paths=source_relative \
         --go-grpc_out=. --go-grpc_opt=paths=source_relative \
         proto/greet.proto
   ```

   This generates:
   - `proto/greet.pb.go` - Message types
   - `proto/greet_grpc.pb.go` - Service interfaces

## 🏃 Running the Project

### Start the Server

In one terminal:

```bash
cd server
go run *.go
```

You should see:
```
sever started at [::]:8080
```

### Run the Client

In another terminal:

```bash
cd client
go run *.go
```

**Note**: Modify `client/main.go` to call different RPC methods:
- `callSayHello(client)` - Unary RPC
- `callSayHelloServerStreaming(client, names)` - Server streaming
- `callSayHelloClientStreaming(client, names)` - Client streaming
- `callSayHelloBidirectionalStreaming(client, names)` - Bidirectional streaming

## 📖 Understanding the RPC Patterns

### 1. Unary RPC (`SayHello`)

**Pattern**: 1 request → 1 response

- **Client**: Sends a single request, waits for a single response
- **Server**: Receives request, processes, returns single response
- **Use Case**: Simple API calls, CRUD operations

**Example Flow**:
```
Client                    Server
  |                         |
  |--- NoParam ------------>|
  |                         | Process
  |<-- HelloResponse -------|
  |                         |
```

### 2. Server Streaming (`SayHelloServerStreaming`)

**Pattern**: 1 request → Multiple responses

- **Client**: Sends one request, receives multiple responses in a loop
- **Server**: Receives request, sends multiple responses using `stream.Send()`
- **Use Case**: Real-time updates, streaming data, progressive results

**Example Flow**:
```
Client                    Server
  |                         |
  |--- NamesList ---------->|
  |<-- HelloResponse 1 -----| stream.Send()
  |<-- HelloResponse 2 -----| stream.Send()
  |<-- HelloResponse 3 -----| stream.Send()
  |<-- EOF -----------------| return nil
```

### 3. Client Streaming (`SayHelloClientStreaming`)

**Pattern**: Multiple requests → 1 response

- **Client**: Sends multiple requests using `stream.Send()`, receives one final response
- **Server**: Receives multiple requests in a loop, returns single response
- **Use Case**: Uploading files, batch processing, aggregating data

**Example Flow**:
```
Client                    Server
  |                         |
  |--- HelloRequest 1 ----->|
  |--- HelloRequest 2 ----->|
  |--- HelloRequest 3 ----->|
  |--- CloseSend() -------->|
  |                         | Process all
  |<-- MessagesList --------|
```

### 4. Bidirectional Streaming (`SayHelloBidirectionalStreaming`)

**Pattern**: Multiple requests ↔ Multiple responses

- **Client**: Uses goroutine to receive while sending in main thread
- **Server**: Receives and sends messages independently
- **Use Case**: Chat applications, real-time collaboration, game servers

**Example Flow**:
```
Client                    Server
  |                         |
  |--- HelloRequest 1 ----->|
  |<-- HelloResponse 1 -----|
  |--- HelloRequest 2 ----->|
  |<-- HelloResponse 2 -----|
  |--- CloseSend() -------->|
  |<-- EOF -----------------|
```

## 🔑 Key Concepts

### Protocol Buffers

Protocol Buffers (protobuf) are Google's language-neutral, platform-neutral mechanism for serializing structured data. The `.proto` file defines:
- **Services**: RPC method definitions
- **Messages**: Data structures (like structs in Go)

### Code Generation

The `protoc` compiler generates:
- **Message types**: Go structs for your data
- **Service interfaces**: Interfaces that your server must implement
- **Client stubs**: Code to call the server from the client

### Goroutines and Channels

Bidirectional streaming requires **goroutines** (lightweight threads) and **channels** (communication pipes) to:
- Send and receive messages simultaneously
- Coordinate between concurrent operations
- Prevent deadlocks

## 🛠️ Development

### Regenerating Proto Files

If you modify `proto/greet.proto`, regenerate the Go code:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       proto/greet.proto
```

### Adding New RPC Methods

1. Add the RPC method to `proto/greet.proto`
2. Regenerate proto files (see above)
3. Implement the method in `server/`
4. Create a client call function in `client/`
5. Update `client/main.go` to call the new method

## 📝 Example Output

### Unary RPC
```
Response from SayHello: Hello
```

### Server Streaming
```
Streaming started
Received: Hello Harsh
Received: Hello Sweta
Streaming finished
```

### Client Streaming
```
Streaming started
Sent request with name: Harsh
Sent request with name: Sweta
Messages: [Hello Harsh Hello Sweta]
Streaming finished
```

### Bidirectional Streaming
```
Bidirectional streaming started
Sent request with name: Harsh
Received: Hello Harsh
Sent request with name: Sweta
Received: Hello Sweta
Bidirectional streaming finished
```

## 🐛 Troubleshooting

### Error: `protoc: command not found`
- Install protoc (see Prerequisites)

### Error: `protoc-gen-go: program not found`
- Run: `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`

### Error: `context deadline exceeded`
- Increase the timeout in client code (e.g., `10*time.Second`)

### Error: `connection refused`
- Make sure the server is running before starting the client

## 📚 Learning Resources

- [gRPC Official Documentation](https://grpc.io/docs/)
- [Protocol Buffers Guide](https://protobuf.dev/)
- [Go gRPC Tutorial](https://grpc.io/docs/languages/go/)

## 📄 License

This project is for educational purposes.

## 👤 Author

Created as a learning project for understanding gRPC patterns in Go.

---

**Happy Learning! 🚀**

