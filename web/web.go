package web

import (
	"context"
	"embed"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	pbs "github.com/joeldejesus1/wasm-go-1/proto/subway"
	log "github.com/sirupsen/logrus"
	"github.com/tarndt/wasmws"
	"google.golang.org/grpc"
)

//go:embed assets
var content embed.FS

type external struct {
	pbs.UnimplementedSubwayMapServer
	ctx context.Context
	wsl *wasmws.WebSockListener
}

type httpExternal struct {
}

type Configuration struct {
	Port uint16
}

// Make a Web server that servers static files and gRPC over websocket.
func Run(
	ctx context.Context,
	config *Configuration,
) error {

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		return err
	}
	// Hi everyone!
	s := grpc.NewServer()

	pbs.RegisterSubwayMapServer(s, e1)

	// server static files
	wsl := wasmws.NewWebSocketListener(ctx)
	e1 := external{ctx: ctx, wsl: wsl}
	httpServer := &http.Server{
		Addr:        fmt.Sprintf(":%d", config.Port),
		Handler:     e1,
		ReadTimeout: 5 * time.Second,
	}

	go func() {
		if err := s.Serve(wsl); err != nil {
			log.Printf("ERROR: Failed to serve gRPC connections; Details: %s", err)
		}
	}()

}

// 1. static files 2. gRPC
func (e1 external) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/grpc":
		e1.wsl.ServeHTTP(w, r)
	default:
		filePathPrefix := "assets"
		// server static files
		log.Infof("serving debug at uri path=%s", r.URL.Path)
		path := strings.TrimPrefix(r.URL.Path, "/")
		if len(path) == 0 || path == "/" {
			path = "/index.html"
		}

		path = strings.TrimPrefix(path, "/")
		file, err := http.FS(content).Open(filePathPrefix + "/" + path)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fileInfo, err := file.Stat()
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		if fileInfo.IsDir() {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		http.ServeContent(w, r, fileInfo.Name(), fileInfo.ModTime(), file)
	}
}
