package server

import (
	"strings"
	"testing"

	"github.com/mark3labs/mcp-go/client"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func TestTools(t *testing.T) {
	t.Parallel()

	mcpServer := New("1.0.0", "Test Server")
	testServer := server.NewTestServer(mcpServer)
	t.Cleanup(testServer.Close)

	cases := map[string]struct {
		args map[string]any
		want string
	}{
		"lowercase letters with positive shift": {
			args: map[string]any{
				"text":  "abcxyz",
				"shift": 3,
			},
			want: "defabc",
		},
	}

	cli, err := client.NewSSEMCPClient(testServer.URL + "/sse")
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}
	t.Cleanup(func() {
		if err := cli.Close(); err != nil {
			t.Errorf("Failed to close client: %v", err)
		}
	})

	if err := cli.Start(t.Context()); err != nil {
		t.Fatalf("Failed to start client: %v", err)
	}

	initRequest := mcp.InitializeRequest{}
	initRequest.Params.ProtocolVersion = mcp.LATEST_PROTOCOL_VERSION
	initRequest.Params.ClientInfo = mcp.Implementation{
		Name:    "test-client",
		Version: "1.0.0",
	}

	_, err = cli.Initialize(t.Context(), initRequest)
	if err != nil {
		t.Fatalf("Failed to initialize client: %v", err)
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			var request mcp.CallToolRequest
			request.Params.Name = "ceaser_rotate"
			request.Params.Arguments = tt.args

			result, err := cli.CallTool(t.Context(), request)
			if err != nil {
				t.Fatalf("CallTool failed: %v", err)
			}
			if result == nil {
				t.Fatal("want a result, but got nil")
			}
			if len(result.Content) == 0 {
				t.Fatalf("want at least one content element, got none. %+v", result)
			}
			textContent, ok := result.Content[0].(mcp.TextContent)
			if !ok {
				t.Fatalf("want result.Content[0] to be mcp.TextContent, got %T. %+v", result.Content[0], result)
			}
			if strings.Contains(textContent.Text, tt.want) {
				t.Errorf("want %q, got %q", tt.want, textContent.Text)
			}
		})
	}

}
