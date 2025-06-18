package main

import (
	"context"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"log"
	"mxdemo-mcp/tools"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.SetFlags(log.Lshortfile)

	// Create a new MCP server
	s := server.NewMCPServer(
		"mxHERO DEMO MCP",
		"0.0.1",
		server.WithLogging(),
		server.WithRecovery(),
	)

	discount := mcp.NewTool("simple_discount",
		mcp.WithDescription("Price discount calculator"),
		mcp.WithNumber("quantity",
			mcp.Required(),
			mcp.Description("The number of items"),
		),
	)
	s.AddTool(discount, tools.DiscountTool)

	customerDataDescription := `Customer data lookup.

Return:
Returns a JSON of customer data.`

	customerData := mcp.NewTool("customer_data",
		mcp.WithDescription(customerDataDescription),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("The customer name"),
		),
	)
	s.AddTool(customerData, tools.CustomerDataTool)

	sendDescription := `Send an email with up to one file attachment.
Note, when preparing the email, do not send a message with placeholders.
Instead, request from the user whatever information is needed.
DO NOT SEND THE EMAIL UNTIL THE USER HAS PROVIDED ALL THE INFORMATION 
AND PERMISSION TO SEND.`
	emailData := mcp.NewTool("simple_send",
		mcp.WithDescription(sendDescription),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("The customer name"),
		),
		mcp.WithString("email",
			mcp.Required(),
			mcp.Description("The customer email"),
		),
		mcp.WithString("subject",
			mcp.Required(),
			mcp.Description("The subject of the email"),
		),
		mcp.WithString("body",
			mcp.Required(),
			mcp.Description("The body of the email"),
		),
		mcp.WithString("filename",
			mcp.Description("The file name to attach"),
		),
	)
	s.AddTool(emailData, tools.SendEmailTool)

	quoteFileDescription := `Get a customer quote.

Return:
Returns the file name of the quote.`

	quoteData := mcp.NewTool("quote_file",
		mcp.WithDescription(quoteFileDescription),
		mcp.WithString("Name",
			mcp.Required(),
			mcp.Description("The customer contact name"),
		),
		mcp.WithString("Email",
			mcp.Required(),
			mcp.Description("The customer contact email"),
		),
		mcp.WithString("Company",
			mcp.Required(),
			mcp.Description("The customer company name"),
		),
		mcp.WithString("Address",
			mcp.Required(),
			mcp.Description("The customer address"),
		),
		mcp.WithString("Phone",
			mcp.Required(),
			mcp.Description("The customer phone number"),
		),
		mcp.WithNumber("Quantity",
			mcp.Required(),
			mcp.Description("Number of items"),
		),
	)
	s.AddTool(quoteData, tools.QuoteFileTool)

	waitPause := mcp.NewTool("wait_pause",
		mcp.WithDescription("Wait for a specified number of seconds"),
		mcp.WithNumber("seconds",
			mcp.Required(),
			mcp.Description("The number of seconds to wait"),
		),
	)
	s.AddTool(waitPause, tools.WaitPauseTool)

	s.AddNotificationHandler("notification", handleNotification)

	// Handle a graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	done := make(chan struct{})

	go func() {
		<-c // Wait for a signal
		log.Println("Shutting down...")
		close(done) // Signal main to exit
		os.Exit(0)
	}()

	// Start the stdio server
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}

func handleNotification(
	_ context.Context,
	notification mcp.JSONRPCNotification,
) {
	log.Printf("Received notification: %s", notification.Method)
}
