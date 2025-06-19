# mxDemo MCP

Assorted mock tools for demo scenarios.

### Arguments

There are only two optional arguments for this MCP. Both are associated with calling DocGen from the Box MCP.

When setting these values, be sure that the user credentials of the Box MCP have access to the Box assets.

* docgen_template_id (optional)
  * Set this argument to hardcode the DocGen template fileID.
* docgen_save_folder_id
  * Set this argument to hardcode the DocGen output folder. 

```json
{
  "mxhero-demo-mcp-server": {
    "args": [
      "-docgen_template_id",
      "<optional - put FileID of DocGen template file>",
      "-docgen_save_folder_id",
      "<optional - put FolderID of DocGen output folder>"
    ],
    "command": "/usr/local/bin/mxdemo"
  }
}
```