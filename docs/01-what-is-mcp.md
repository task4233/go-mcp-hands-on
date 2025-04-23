# MCPの概要

今回のハンズオンの中心のコンテンツではないので、深く立ち入りません。

詳しくは、以下の公式ドキュメントか、他の方が書かれた入門記事をご覧ください。

https://modelcontextprotocol.io/introduction

### MCPとは

- Model Context Protocolの略
- LLM に与えるコンテキストを標準化するオープンプロトコル
  - AIアプリの「USB-Cポート」のように多様なデータ源やツールを統一的に接続できるようにすることを目的とする
    - ※ この「USB-Cポート」という表現は、MCPの公式ドキュメントに実際に登場している

### Architecture  

- ホスト・クライアント・サーバの3層モデルで、ホストが複数クライアントを束ねて各サーバに接続する
  - ホストの例: Claude Desktop
  - クライアントの例: Claude Desktopに内包
  - サーバーの例: notion-mcp, playwright-mcp etc...

### Base Protocol  

- JSON-RPC 2.0ベースのプロトコル
  - ライブラリを使ってMCPサーバーを実装する分には、基本的に意識することはない
- トランスポートのレイヤーはBase Protocolと分かれている

### MCPサーバーが提供できるもの

- Resources
  - https://modelcontextprotocol.io/docs/concepts/resources
- Prompts
  - https://modelcontextprotocol.io/docs/concepts/prompts
- Tools
  - https://modelcontextprotocol.io/docs/concepts/tools
  - 今回扱うのはこれ

### Tools

- LLM が外部システムを呼び出す関数を標準化したもの
  - `tools/list` で発見し `tools/call` のRPCを実行。各ツールは `name`,`description`, `inputSchema` を必須で持つ
- 結果はテキスト・画像・リソース埋め込みなど複数形式で返送し、失敗時は結果の `isError` フィールドを `true` に設定するか、JSON-RPCエラーを返す
- `human-in-the-loop` が推奨される
  - プロトコルとしてはAIモデルが自動的にツールを検出して使うことができ、ユーザー対話形式であることが強制されないが、UIで確認プロンプトを挟み入力検証・権限制御を行うことで、誤操作や悪用を防ぐことが強く推奨される

### Transport

Transportの形式は2つ

- Stdio (標準入出力)
- Streamable HTTP
  - 2025-03-26仕様より前はSSE (Server Sent Events) だった
    - 2025年4月時点ではまだ多くのクライアント / サーバー実装にSSEが残っている
  - こちらが参考になります https://azukiazusa.dev/blog/mcp-server-streamable-http-transport/

ローカル環境で動かす場合は主にStdioが使われる。(2025年4月時点ではClaude DesktopはStdioしかサポートしていない)

#### MCPサーバーの設定例

多くのMCPホストが、以下のような形式で設定を行う。

```json
"servers": {
  "local-server": {
    // "type": "stdio", // VS Codeでは必要
    "command": "npx",
    "args": ["-y", "@syumai/example-mcp-server"]
  },
  "remote-server": {
    // "type": "sse", // VS Codeでは必要
    "url": "http://mcp.syum.ai/sse",
    "headers": { "VERSION": "1.0" }
  }
}
```

例えば、 `command` に指定された `npx` コマンドは、ホストが起動する度に実行される。
クライアントとサーバーは標準入出力でコミュニケーションするため、よくHTTPサーバーを起動する時のように *どのポートが使われているか* といったことを気にする必要がない。

以下はVS Codeのドキュメントの「[Configuration format](https://code.visualstudio.com/docs/copilot/chat/mcp-servers#_configuration-format)」セクションを抜粋し、日本語化したものです (イメージしやすくするために掲載しています)

---

#### 設定フォーマット

- `"servers": {}` フィールドには MCP サーバーの一覧を指定し、Claude Desktop の設定フォーマットに準拠します。  
- サーバー名をキーにして、サーバー設定を値として指定します。VS Code の MCP サーバー一覧にサーバー名が表示されます。  
- `"inputs": []` フィールドでは、設定値のプレースホルダーを定義できます。サーバー起動時に初回のみ VS Code が入力を求め、安全に保存します。入力値を非表示にするには `password: true` を設定します。
- 事前定義変数も利用可能です。
  - https://code.visualstudio.com/docs/reference/variables-reference

##### Stdio接続

| フィールド  | 説明                             | 例                                                      |
| ----------- | -------------------------------- | ------------------------------------------------------- |
| `type`      | サーバー接続タイプ               | `"stdio"`                                              |
| `command`   | サーバー実行コマンド             | `"npx"`, `"node"`, `"python"`, `"docker"`               |
| `args`      | コマンドに渡す引数の配列         | `["server.py", "--port", "3000"]`                      |
| `env`       | サーバー用環境変数               | `{"API_KEY": "${input:api-key}"}`                       |
| `envFile`   | 追加の環境変数を読み込む `.env` へのパス | `"${workspaceFolder}/.env"`                             |

##### SSE接続

| フィールド  | 説明                              | 例                                      |
| ----------- | --------------------------------- | --------------------------------------- |
| `type`      | サーバー接続タイプ                | `"sse"`                                 |
| `url`       | サーバーの URL                   | `"http://localhost:3000"`              |
| `headers`   | サーバーへの HTTP ヘッダー        | `{"API_KEY": "${input:api-key}"}`      |
