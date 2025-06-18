package tools

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type MailAttachment struct {
	Name string
	Data []byte
}

type SimpleMail struct {
	To             string `json:"to" jsonschema:"required,description=The recipient of the email"`
	Subject        string `json:"subject" jsonschema:"required,description=The subject of the email"`
	Body           string `json:"body" jsonschema:"required,description=The body of the email"`
	AttachmentName string `json:"attachmentName" jsonschema:"description=The file name of the attachment"`
}

func SendEmailTool(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultError("this tool is not implemented"), nil
}
