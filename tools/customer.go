package tools

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/mark3labs/mcp-go/mcp"
	"strings"
)

// CustomerDataTool returns hard-coded customer data
func CustomerDataTool(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// not using the parameter
	nameArg, nameErr := request.RequireString("name")
	if nameErr != nil {
		return nil, nameErr
	}
	name := strings.TrimSpace(nameArg)
	if name == "" {
		return nil, errors.New("name is required")
	}

	// Simulate a CRM lookup
	dataMap := map[string]string{
		"Name":    name,
		"Email":   "alex@test.mxhero.com",
		"Company": "Acme Inc.",
		"Address": "123 Main St, Anytown, USA",
		"Phone":   "555-456-7890",
	}

	jsonString, err := json.Marshal(dataMap)
	if err != nil {
		return nil, err
	}

	return mcp.NewToolResultText(string(jsonString)), nil
}
