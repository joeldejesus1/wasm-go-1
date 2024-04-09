# Demonstration of Go with WASM and gRPC

We create a simple website that lets a user ask for directions in a simple subway system.  We use Go to host a webserver that serves a WASM over http and gRPC over websockets.

# Build

Ensure that you have the minimum required version of Golang installed as indicated in the go.mod file, and that GOROOT and GOPATH are correctly set in your env.

```bash
./contrib/proto.sh
make
```

# Run

```bash
./bin/web.server
```
Browse the website at localhost:8080 .  
