package simplemcp

import (
	"github.com/mark3labs/mcp-go/server"
	"github.com/rickKoch/go-mcp/pkg/hn"
)

func NewServer() *server.MCPServer {
	client := hn.NewClient()
	s := server.NewMCPServer("simple-mcp-server-example", "0.0.0")
	s.AddTool(GetNewStories(client))
	s.AddTool(GetTopStories(client))

	return s
}
