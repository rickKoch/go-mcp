package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"
	"syscall"

	mcpserver "github.com/mark3labs/mcp-go/server"
	simplemcp "github.com/rickKoch/go-mcp/pkg/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	server := simplemcp.NewServer()
	stdioServer := mcpserver.NewStdioServer(server)

	errC := make(chan error, 1)
	go func() {
		errC <- stdioServer.Listen(ctx, io.Reader(os.Stdin), io.Writer(os.Stdout))
	}()

	select {
	case <-ctx.Done():
		_, _ = fmt.Fprintf(os.Stderr, "shutting down server...")
	case err := <-errC:
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}

}
