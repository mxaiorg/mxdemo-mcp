# mxDemo MCP

Assorted mock tools for demo scenarios.

## Compiling

```bash
make
cp bin/mxdemo <target folder>
```

Be sure to note the full path of the target file location. The full path needs to be put in the configuration (see below)  "command" field.

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

Replace "/usr/local/bin/mxdemo" with the location of your compiled binary.