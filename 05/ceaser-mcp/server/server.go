package server

import (
	"context"
	"encoding/json"
	"errors"
	"log/slog"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/syumai/go-mcp-hands-on/05/ceaser-mcp/ceaser"
)

func New(version, description string) *server.MCPServer {
	s := server.NewMCPServer(description, version)

	tool := mcp.NewTool("ceaser_rotate",
		mcp.WithDescription("Rotate a string by a given number of positions. It is used to encrypt or decrypt text of Ceaser Cipher."),
		mcp.WithString("text",
			mcp.Required(),
			mcp.Description("Text to be rotated"),
		),
	)

	s.AddTool(tool, rotateHandler)

	return s
}

const errRequiredText = "text must be a string"

type ShiftResult struct {
	Shift int    `json:"shift"`
	Text  string `json:"text"`
}

func rotateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	text, ok := request.Params.Arguments["text"].(string)
	if !ok {
		return mcp.NewToolResultError(errRequiredText), errors.New(errRequiredText)
	}

	// Create results for all shifts from 1 to 25
	var results []ShiftResult

	// Add original text as shift 0
	results = append(results, ShiftResult{
		Shift: 0,
		Text:  text,
	})

	for shift := 1; shift <= 25; shift++ {
		rotatedText := ceaser.RotN(text, shift)
		results = append(results, ShiftResult{
			Shift: shift,
			Text:  rotatedText,
		})
	}

	// Convert results to JSON
	jsonData, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return mcp.NewToolResultError("Failed to marshal results to JSON"), err
	}

	slog.Info("Rotate handler called", "text", text)

	return mcp.NewToolResultText(string(jsonData)), nil
}
