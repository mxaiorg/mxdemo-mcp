# Prebuilt Binaries

This folder contains pre-built binaries to facilitate deployment.

## Mac ARM (Apple Silicon—M1, M2 ...) 

**mxdemo-mac-arm-installer.pkg**

* Double-click the package to install.
* md5sum
  * 9eee519b879164ed78a4265956028075  prebuilt/mxdemo-mac-arm-installer.pkg

## Mac AMD (Apple Intel)

**mxdemo-mac-intel-installer.pkg**

* Double-click the package to install.
* md5sum
  * b9630e1da3221b0a14839eec8c63da04  prebuilt/mxdemo-mac-intel-installer.pkg

## Configuration

### Mac

The Mac package installs mxdemo-mac-arm/intel to your `/usr/local/bin` directory. As such, set the command field in your configuration to:

* Mac ARM: `/usr/local/bin/mxdemo-mac-arm`
* Mac Intel: `/usr/local/bin/mxdemo-mac-intel`

For example:

```json
{
  "mcpServers": {
    "mxhero-mcp-server": {
      "command": "/usr/local/bin/mxdemo-mac-arm",
      "args": [
        ... see arguments in repo readme ...
      ]
    }
  }
}
```