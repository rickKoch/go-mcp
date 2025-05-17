# Model Context Protocol (MCP) example

![example](/data/go-mcp3.gif "example")

This is a simple project that demonstrates how to create an MCP (Model Context
Protocol) server in Golang and use it within VS Code Copilot.

It was built as part of a blog post explaining the basics of MCP and how it can
be used to connect tools and data sources with AI-powered environments.

📖 Read the blog post:     
[Model Context Protocol(MCP)](https://rickkoch.github.io/posts/model-context-protocol).

## 🚀 What It Does

* Implements a basic MCP server in Go.
* Exposes one or more tools that can be used by AI systems supporting MCP.
* Demonstrates integration with VS Code Copilot via MCP.

## 🛠 Technologies Used

* Go (Golang)
* MCP Protocol (Open Standard)
* VS Code Copilot (with MCP support)

## 📦 How to Use

### 1. Clone the repository:
```bash
git clone https://github.com/rickKoch/go-mcp.git
cd go-mcp
```

### 2. Build the docker image:
```bash
./build.sh
```

### 3. Copy the following configuration at the bottom of the VS Code user settings file(`settings.json`):
```json
"mcp": {
    "servers": {
        "simplemcp": {
        "command": "docker",
        "args": ["run", "-i", "--rm", "simplemcp"]
        },
    }
}
```

### 4. Once the server is running you should toggle the Agent and try it out