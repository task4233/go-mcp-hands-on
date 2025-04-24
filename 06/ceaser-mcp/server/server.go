package server

import "github.com/mark3labs/mcp-go/server"

func New() *server.MCPServer {
	s := server.NewMCPServer(
		"Ceaser Cipher",
		"1.0.0",
	)
	// Add tool here
	// name: "ceaser_rotate"
	// description: "Rotate a string by a given number of positions. It is used to encrypt or decrypt text of Ceaser Cipher."
	return s
}

func rotateHandler() {
	// TODO: implement
}
