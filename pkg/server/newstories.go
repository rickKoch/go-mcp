package simplemcp

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/rickKoch/go-mcp/pkg/hn"
)

type newStoriesClient interface {
	NewStories(number int) ([]hn.Story, error)
}

func GetNewStories(client newStoriesClient) (mcp.Tool, server.ToolHandlerFunc) {
	readOnlyHint := true
	tool := mcp.NewTool("get_hn_new_stories", mcp.WithDescription("Fetch the latest stories from Hacker News. You should specify the number of stories to retrieve, up to a maximum of 500."),
		mcp.WithToolAnnotation(mcp.ToolAnnotation{
			Title:        "Get new Hacker News stories",
			ReadOnlyHint: &readOnlyHint,
		}),
		mcp.WithNumber("number", mcp.Required(), mcp.Description("The number of new stories to retrieve from Hacker News (maximum 500).")),
	)

	handler := func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		// fmt.Fprintf(os.Stderr, "value: %#v, type: %T", request.Params.Arguments["number"], request.Params.Arguments["number"])
		if _, ok := request.Params.Arguments["number"]; !ok {
			return mcp.NewToolResultError("missing required parameter `number`"), nil
		}

		if _, ok := request.Params.Arguments["number"].(float64); !ok {
			return mcp.NewToolResultError("parameter `number` is not of type int"), nil
		}

		result, err := client.NewStories(int(request.Params.Arguments["number"].(float64)))
		if err != nil {
			return nil, fmt.Errorf("failed to get Hacker News new stories: %w", err)
		}

		r, err := json.Marshal(result)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal response: %w", err)
		}

		return mcp.NewToolResultText(string(r)), nil
	}

	return tool, handler
}
