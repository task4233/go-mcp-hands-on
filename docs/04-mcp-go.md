# mcp-goの基本

READMEのサンプルを確認してみましょう。

```go
package main

import (
    "context"
    "errors"
    "fmt"

    "github.com/mark3labs/mcp-go/mcp"
    "github.com/mark3labs/mcp-go/server"
)

func main() {
    // MCPサーバーを作成
    s := server.NewMCPServer(
        "Demo 🚀",
        "1.0.0",
    )

    // ツールを追加
    tool := mcp.NewTool("hello_world",
        mcp.WithDescription("Say hello to someone"),
        mcp.WithString("name",
            mcp.Required(),
            mcp.Description("Name of the person to greet"),
        ),
    )

    // ツールハンドラーを追加
    s.AddTool(tool, helloHandler)

    // Stdioサーバーを起動
    if err := server.ServeStdio(s); err != nil {
        fmt.Printf("Server error: %v\n", err)
    }
}

func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
    name, ok := request.Params.Arguments["name"].(string)
    if !ok {
        return nil, errors.New("name must be a string")
    }

    return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}
```

https://github.com/mark3labs/mcp-go/blob/v0.23.1/README.md

まずは、実際に動作する様子を確認するため、MCPホストの設定にこのMCPサーバーを登録してみましょう。

VS Codeをご利用の場合は、以下の手順を実行します。

1. `Cmd + Shift +P` からコマンドパレットを開く
2. `MCP: Add Server` を実行
3. `Command: (stdio)` を選択
4. `Enter Command` と聞かれるので `go` と入力
5. `Enter Server ID` と聞かれるので `hello` と入力
6. 表示されるJSONの `servers.hello` を以下の内容に修正する

```json
"servers": {
  "hello": {
    "type": "stdio",
    "command": "go",
    "args": [
      "run",
      "github.com/syumai/go-mcp-hands-on/examples/hello@latest"
    ]
  }
}
```

ここまで設定が完了したら、 `hello` のキーの上に表示される `Start` をクリックして、GitHub CopilotのAgentを開いてください。

チャットに「syumaiにhelloと言ってください」と入力すると、 `hello_world` のツールが選択され、呼ばれることが確認できます。

## mcp-goの概要

ここではざっくりと構造の説明を行います。

詳細については[Go Doc](https://pkg.go.dev/github.com/mark3labs/mcp-go@v0.23.1)を参照してください。

### 基本構造

- `mcp` packageは、Server / Clientに依存しないToolなどの型定義が置かれている
- `server` packageには、Serverの実装が置かれている
- [server.MCPServer](https://github.com/mark3labs/mcp-go/blob/v0.23.1/server/server.go#L143) 型は、特定のTransportに依存しないMCP Serverの実装を提供する
  - これを、Stdio / SSEのTransportでWrapして使う

### Tool定義の構造

- `mcp.NewTool` でToolを定義
- `mcp.WithDescription` でToolの説明を登録
- パラメーターの定義はコードで行う
  - `mcp.WithString` で文字列引数を定義
    - `mcp.Required` で引数を必須化
    - `mcp.Description` で引数の説明を登録
- 結果の形式については特に指定する必要がない

```go
tool := mcp.NewTool("hello_world",
    mcp.WithDescription("Say hello to someone"),
    mcp.WithString("name",
        mcp.Required(),
        mcp.Description("Name of the person to greet"),
    ),
)
```

### Toolハンドラーの構造

- シグニチャの構成
  - パラメーター: `context.Context` と `mcp.CallToolRequest`
  - 戻り値: `mcp.CallToolResult` と `error`
- 引数は、`mcp.CallToolRequest` から取得
  - 引数には自動的に型が付かないため、`.(string)` で型アサーションする
- 結果は、 `mcp.NewToolResultText` などで返す
  - 画像なども返せる

```go
func helloHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
    name, ok := request.Params.Arguments["name"].(string)
    if !ok {
        return nil, errors.New("name must be a string")
    }

    return mcp.NewToolResultText(fmt.Sprintf("Hello, %s!", name)), nil
}
```