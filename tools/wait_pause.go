package tools

import (
	"context"
	"errors"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"time"
)

func WaitPauseTool(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	seconds := mcp.ParseInt(request, "seconds", 1)
	if seconds == 0 {
		return nil, errors.New("query must be a positive number")
	}

	time.Sleep(time.Duration(seconds) * time.Second)

	response := fmt.Sprintf("Paused for %d seconds", seconds)

	return mcp.NewToolResultText(response), nil
}
