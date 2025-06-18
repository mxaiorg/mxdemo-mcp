package tools

import (
	"context"
	"errors"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
)

func DiscountTool(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	quantity := mcp.ParseInt(request, "quantity", 0)
	if quantity == 0 {
		return nil, errors.New("query must be a positive number")
	}

	discount := discountCalc(quantity)

	response := fmt.Sprintf("Discount: %d", discount)

	return mcp.NewToolResultText(response), nil
}

func discountCalc(quantity int) int {
	var discount int

	switch {
	case quantity >= 2000:
		discount = 20
	case quantity >= 1000:
		discount = 10
	default:
		discount = 0
	}
	return discount
}
