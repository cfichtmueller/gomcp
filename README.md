# gomcp

[Go 1.24](https://golang.org/) | [MIT License](LICENSE)

A Go library for building **Model Context Protocol (MCP)** servers. This library provides a simple and intuitive way to create MCP servers that can expose tools and resources to AI models.

> **⚠️ Early Development**: This library is in very early stages of development. The API is not yet stable and may change significantly. Nothing is set in stone yet, and it's not yet fully compliant with the MCP specification.

## What is MCP?

The Model Context Protocol (MCP) is a standard for connecting AI models to external tools and data sources. It enables AI assistants to interact with various services, databases, APIs, and other resources through a standardized interface.

## Features

- 🚀 **Simple API**: Easy-to-use Go API for creating MCP servers
- 🛠️ **Tool Support**: Define and expose tools that AI models can call
- 🌐 **HTTP Transport**: Built-in HTTP transport for web-based integrations

## Examples

Check out the `examples/` directory for complete working examples:

- [`examples/hello_server/`](examples/hello_server/) - Basic server with simple tools including hello world and calculator functions
- More examples coming soon...

## Development Status

This library is currently in **early development**. Here's what's implemented:

✅ **Core Features:**
- Basic MCP server implementation
- Tool definition and execution
- HTTP transport
- Basic protocol handling

🚧 **In Progress:**
- Resource support
- Better error handling

📋 **Planned:**
- Full MCP specification compliance
- Authentication support

## Contributing

This project doesn't accept contributions yet.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Built for the [Model Context Protocol](https://modelcontextprotocol.io/) specification
- Inspired by the need for better AI-tool integration in Go applications

---

**Note**: This library is not yet production-ready. Use at your own risk and expect breaking changes in future versions.
