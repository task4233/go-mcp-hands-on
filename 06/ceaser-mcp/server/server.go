package server

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/syumai/go-mcp-hands-on/05/ceaser-mcp/ceaser"
)

type Server struct {
	*server.MCPServer
	timeout time.Duration
}

func New(version, description string) *Server {
	slog.Info("Ceaser Cipher Server", "version", version, "description", description)

	ss := server.NewMCPServer(description, version)

	tool := mcp.NewTool("ceaser_rotate",
		mcp.WithDescription("Rotate a string by a given number of positions. It is used to encrypt or decrypt text of Ceaser Cipher."),
		mcp.WithString("text",
			mcp.Required(),
			mcp.Description("Text to be rotated"),
		),
		mcp.WithNumber("shift",
			mcp.Required(),
			mcp.Description("Number of positions to rotate the text"),
		),
	)

	s := &Server{
		MCPServer: ss,
		timeout:   5 * time.Second,
	}

	ss.AddTool(tool, s.rotateHandler)

	return s
}

func (s *Server) rotateHandler(ctx context.Context, request mcp.CallToolRequest) (res *mcp.CallToolResult, e error) {
	slog.Info("rotateHandler called", "request", request)

	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	req, err := parseRotateRequest(request)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), err
	}

	result := ceaser.RotN(req.Text, req.Shift)
	return mcp.NewToolResultText(result), nil
}

type rotateRequest struct {
	Text  string
	Shift int
}

func parseRotateRequest(request mcp.CallToolRequest) (*rotateRequest, error) {
	text, ok := request.Params.Arguments["text"].(string)
	if !ok {
		return nil, errors.New("text must be a string")
	}

	// Handle shift parameter as float64 (JSON number) and convert to int
	shiftFloat, ok := request.Params.Arguments["shift"].(float64)
	if !ok {
		return nil, errors.New("shift must be a number")
	}
	shift := int(shiftFloat)

	return &rotateRequest{
		Text:  text,
		Shift: shift,
	}, nil
}
