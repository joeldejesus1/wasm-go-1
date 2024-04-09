MAKE := make

GO_BIN := go


ifeq ($(DEBUG),)
# had garble, but it does not work with 
	GOFLAGS :=  
	LDFLAGS := -s -w 
else
	GOFLAGS :=
	LDFLAGS :=
endif

ifeq ($(uname -o | awk '{print tolower($0)}'), "linux")
	OS := linux
else
	OS := darwin
endif


all: build

.PHONY: all setup wasm

build: setup wasm server

server: setup wasm ./cmd/main.go
	GOOS=$(OS) GOARCH=amd64 $(GO_BIN) build $(GOFLAGS) -ldflags="$(LDFLAGS)" -mod=readonly -buildvcs=false -o bin/web.server github.com/joeldejesus1/wasm-go-1/cmd

wasm: setup ./wasm/main.go
	GOOS=js GOARCH=wasm $(GO_BIN) build $(GOFLAGS) -ldflags="$(LDFLAGS)" -mod=readonly -buildvcs=false -o web/assets/sw.wasm github.com/joeldejesus1/wasm-go-1/wasm

setup: ./go.mod
	$(GO_BIN) mod tidy
	mkdir ./bin || true
	@bash ./contrib/proto.sh go

clean:
	rm web/assets/*.wasm 2>/dev/null || true
