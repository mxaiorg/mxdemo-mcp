package tools

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/mark3labs/mcp-go/mcp"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strconv"
	"text/template"
)

var (
	DocgenTmplId       string
	DocgenSaveFolderId string
)

var quoteFileStepInstructs = `Follow these steps to download a quote file:

# Step 1:

Call box_docgen_create_batch_tool with the following parameters:
{{if .FileId}}
file_id: {{.FileId}}
{{- else}}
Work with the user to get the file id of the DocGen template.
Once identified, use it as the 'file_id' parameter value.
{{- end}}
{{if .FolderId}}
destination_folder_id: {{.FolderId}}
{{- else}}
Work with the user to get the folder id of the DocGen output folder.
Once identified, use it as the 'destination_folder_id' parameter value.
{{end}}
user_input_file_path: /tmp/quote.json

With the returned batch id proceed to step 2. 

# Step 2:

Call box_docgen_list_jobs_by_batch_tool with the following parameters:

batch_id: <batch id from step 1>

Check if the returned job list contains "status": "completed". If so, check the "output_file" 
field and get the file id. If "status" is not "completed", call wait_pause for 5 seconds.
When wait_pause returns, repeat this step.

DO NOT repeat this step (Step 2) more than 3 times. If repeating more than 3 times, respond with 
the error message and do NOT proceed to step 3, and do NOT attempt to download a file.

# Step 3:

Call the box_download_file_tool. With the following parameters:
- file_id: <file id from step 2>
- save_file: true
- save_path: /tmp/quote.pdf

If there is an error, respond with 'ERROR' and the error message.
Otherwise, respond with 'OK'.
`

// QuoteFileTool takes a map of key values, does some pricing calculations
// and saves everything to a json file. Returns detailed instructions on
// how to generate and save a quote file.
func QuoteFileTool(_ context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	unitPrice := 10

	jsonMap := make(map[string]string)

	titleCaser := cases.Title(language.English)
	for key := range request.GetArguments() {
		val := mcp.ParseString(request, key, "")
		if val != "" {
			upperArg := titleCaser.String(key)
			jsonMap[upperArg] = val
		}
	}

	// CALCULATE

	// Quantity
	var quantity int
	var err error
	if quantityStr, ok := jsonMap["Quantity"]; ok {
		quantity, err = strconv.Atoi(quantityStr)
		if err != nil {
			return nil, err
		}
		jsonMap["Quantity"] = strconv.Itoa(quantity)
	}

	amount := quantity * unitPrice
	jsonMap["Amount"] = fmt.Sprintf("%d", amount)

	discountPct := discountCalc(quantity)
	jsonMap["DiscountPct"] = fmt.Sprintf("%d", discountPct)

	discountAmt := discountPct * unitPrice * quantity
	jsonMap["DiscountAmt"] = fmt.Sprintf("%d", discountAmt)

	subtotal := amount - discountAmt
	jsonMap["Subtotal"] = fmt.Sprintf("%d", subtotal)

	taxAmt := float64(subtotal) * 0.05
	jsonMap["TaxAmt"] = fmt.Sprintf("%.2f", taxAmt)

	total := float64(subtotal) - taxAmt
	jsonMap["Total"] = fmt.Sprintf("%.2f", total)

	jsonBytes, err := json.MarshalIndent(jsonMap, "", "    ")
	if err != nil {
		return nil, err
	}

	// write out json that will be used for DocGen call
	err = os.WriteFile("/tmp/quote.json", jsonBytes, 0644)
	if err != nil {
		return nil, err
	}

	tmplMap := map[string]string{
		"FileId":   DocgenTmplId,
		"FolderId": DocgenSaveFolderId,
	}

	quoteInstructions, tmplErr := fillQuoteTmpl(quoteFileStepInstructs, tmplMap)
	if tmplErr != nil {
		return nil, tmplErr
	}

	return mcp.NewToolResultText(quoteInstructions), nil
}

func fillQuoteTmpl(tmpl string, jsonMap map[string]string) (string, error) {
	t, err := template.New("quote").Parse(tmpl)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if execErr := t.Execute(&buf, jsonMap); execErr != nil {
		return "", execErr
	}

	return buf.String(), nil
}
