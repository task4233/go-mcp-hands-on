package main

import (
	"fmt"
	"log/slog"

	"github.com/mark3labs/mcp-go/server"
	ceaserserver "github.com/syumai/go-mcp-hands-on/05/ceaser-mcp/server"
)

const (
	version     = "1.0.0"
	description = "Ceaser Cipher"
)

func main() {
	s := ceaserserver.New(version, description)
	slog.Info("Ceaser Cipher Server", "version", version, "description", description)
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
