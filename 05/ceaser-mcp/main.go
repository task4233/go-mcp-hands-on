package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	ceaserserver "github.com/syumai/go-mcp-hands-on/05/ceaser-mcp/server"
)

func main() {
	s := ceaserserver.New()
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
