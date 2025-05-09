package main

import (
	"context"
	"os/exec"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func main() {
	// Create a new MCP server
	s := server.NewMCPServer("EmacsLispServer", "1.0.0")

	// Define the execute_lisp tool
	executeLispTool := mcp.NewTool("execute_lisp",
		mcp.WithDescription("Executes a Lisp command in Emacs"),
		mcp.WithString("command", mcp.Description("the lisp command to run"), mcp.Required()),
	)

	// Define the handler for the tool
	executeLispHandler := func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		command := mcp.ParseString(request, "command", "emacs-version")
		cmd := exec.Command("emacsclient", "--eval", command)
		output, err := cmd.CombinedOutput()
		if err != nil {
			return nil, err
		}
		result := string(output)

		return &mcp.CallToolResult{
			Result: mcp.Result{
				Meta: map[string]interface{}{},
			},
			Content: []mcp.Content{
				mcp.NewTextContent(result),
			},
			IsError: false,
		}, nil
	}

	// Add the tool to the server
	s.AddTool(executeLispTool, executeLispHandler)

	// Start the server
	server.ServeStdio(s)
}
