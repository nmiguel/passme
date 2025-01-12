# Passme

## About this tool

Passme is a tool for copying and managing tokens or passwords. It is not secure and does not encrypt the data. This may be implemented in the future, but if security is a concern, you should not use this tool.

This tool is designed for use when prompted for a password or token. Simply background the current task or use a separate terminal tab or split to quickly get the token into your clipboard, ready for pasting.

## Installation

To install, use Go's `install` command:

`go install github.com/nmiguel/passme@latest`

### Dependencies

- Go version 1.20 or higher
- `xclip` or `xsel` for clipboard support

## Usage

To copy a token or password to your clipboard:
```sh
passme c <token_name>
```

To open the TUI:
```sh
passme
```
## License

Passme is licensed under the GNU GPL 3.0 License. See the `LICENSE` file for details.
