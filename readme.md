# mxDemo MCP

Assorted mock tools for demo scenarios.

**Why Go for MCP deployment**

Unlike Python or Javascript MCPs, Go compiles to native static binary. Once compiled for a target architecture (e.g., Mac ARM, Windows Intel) and installed, no additional dependencies (software) are required on the user's device.

## Quickstart

1. To get started quickly, see the prebuilt binaries section below for your operating system and machine architecture.
2. Follow the installation instructions.

## Requirements

- GO 1.22 or higher ([download](https://go.dev/doc/install))
- mxHERO Vector Search credentials (token)
  - A demo token can be obtained at https://lab4-api.mxhero.com/demo_signup
  - For production tokens, contact mxHERO.

## Installation
1. Clone the repository
```sh
git clone https://github.com/mxaiorg/mxdemo-mcp
cd mxdemo-mcp
go mod tidy
```

2. Compile

Be sure to compile for the architecture of the user's computer. You will need to match the operating system and processor architecture. The included Makefile provides for a few of the most common.

| OS      | Architecture              | Make command       |
|---------|---------------------------|--------------------|
| Windows | Intel                     | make windows-intel |
| Mac     | Arm (Mac silicon - M1...) | make mac-arm       |
| Mac     | Intel                     | make mac-intel     |
| Linux   | Intel                     | make linux-intel   |

For more operating systems and architectures see Go compilation documentation.

**Example build**

```shell
make mac-arm
```

After `make` is run it will place the program (binary) in the `bin`folder. Copy this binary to the user's computer and see the configuration instructions below.

For example: `cp bin/mxmcp-mac-arm ~/user`

**Note**
* Some platforms, like MacOS, will require additional permissions before allowing the program to be run on another machine.


## Prebuilt Binaries
For convenience the `prebuilt` folder contains prebuilt binaries and signed installation packages. See the "readme" file in that folder for more information.

## Installation

If **NOT** installing with an installation package (prebuilt), do the following:

1. Copy the binary (of matching operating systems and architecture) to the user's computer. Place the file somewhere the user has permission to access. For example, the user's home directory.


2. Ensure the user has permission to run the program (execute)
  - For example, on Mac & Linux `chmod 755 mxdemo`

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

### Note about DocGen demo

The "create quote" demonstration uses a predefined template with fields the demo MCP will supply values for. You can find a copy of this template in the `assets` folder.

This file should be the target template file of the quote_file tool. You should upload this file into Box as a DocGen template.