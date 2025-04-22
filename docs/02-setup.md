# 環境構築

## MCPクライアントのセットアップ

MCPサーバーを開発して動かすには、MCPクライアントが必要となりますので、何かしらのクライアントを事前にインストールおよび、起動しておいてください。

基本的には、MCPサーバーとの連携がこなれている[Claude Desktop](https://claude.ai/download)をおすすめします。
2025年4月時点で、筆者が最もよく使っているAIエージェントはClaude Desktopとなっています。

Claude DesktopにおけるMCPサーバーの設定は以下のリンクを参考に行えます。
https://modelcontextprotocol.io/quickstart/user

本ハンズオンではClaude Desktopの使用を前提としますが、以下のツールに関してはサポート可能です。

### エディター系

* [VS Code (GitHub Copilot)](https://code.visualstudio.com/)
  - https://code.visualstudio.com/docs/copilot/chat/mcp-servers

### VS Code拡張系

* [Cline](https://github.com/cline/cline)
* [Roo Code](https://github.com/RooVetGit/Roo-Code)

### CLI tool系

* [Claude Code](https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/overview)
  - https://docs.anthropic.com/en/docs/agents-and-tools/claude-code/tutorials#set-up-model-context-protocol-mcp
* [oterm](https://github.com/ggozad/oterm)

以下のエディターでもMCPサーバーとの連携が可能ですが、筆者はまだ設定を行ったことがないため、ご自身でドキュメントを参考に作業を行ってください。　

* [Cursor](https://github.com/cursor-ai/cursor)
  - https://docs.cursor.com/context/model-context-protocol
* [Windsurf](https://github.com/windsurf-io/windsurf)
  - https://docs.windsurf.com/windsurf/mcp

## 開発環境のセットアップ

Goさえインストールしてあれば、エディタは何でも構いません。ご自身の好きなエディタをご利用ください。
