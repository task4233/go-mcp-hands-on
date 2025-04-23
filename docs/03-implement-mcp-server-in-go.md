# Go言語によるMCPサーバーの実装についての概要

## 公式SDKについて

2025年4月現在、MCPの公式のSDKが存在するのは[以下の言語](https://github.com/modelcontextprotocol?q=sdk&type=all&language=&sort=)です。

- Java
- Python
- Kotlin
- TypeScript
- C#
- Swift
- Rust

見ての通り、Go言語の公式のSDKは存在していないため、サードパーティーのライブラリを使用するのが一般的となります。

## サードパーティーのライブラリ

主要なサードパーティーのライブラリとしては以下のものが存在します。

- [github.com/mark3labs/mcp-go](https://github.com/mark3labs/mcp-go)
- [github.com/metoro-io/mcp-golang](https://github.com/metoro-io/mcp-golang)
- [github.com/ktr0731/go-mcp](https://github.com/ktr0731/go-mcp)

2025年4月現在は、 `mark3labs/mcp-go` (以降、単にmcp-goと呼びます) の採用例が最も多いです。
[github-mcp-server](https://github.com/github/github-mcp-server)もこのライブラリを使用して実装されています。
初めの一歩としては、最も使われているライブラリを使うのが適切と思いますので、本ハンズオンでは、このライブラリを使用してMCPサーバーを実装していきます。

ただし、 `mark3labs/mcp-go` には型安全性が十分でない点が多々存在しているので、その辺りが気になる方は、自動コード生成ツールを内包している `ktr0731/go-mcp` を試してみてみるとよいかもしれません。
