package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	ceaserserver "github.com/syumai/go-mcp-hands-on/05/ceaser-mcp/server"
)

const (
	version     = "1.0.0"
	description = "Ceaser Cipher(06)"
)

func main() {
	s := ceaserserver.New(version, description)
	if err := server.ServeStdio(s.MCPServer); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
