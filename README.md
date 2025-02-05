# Passme

## About this tool

Passme is a tool for copying and managing tokens or passwords. It is not secure and does not encrypt the data (but the naked token is never shown). This may be implemented in the future, but if security is a concern, you should not use this tool.

This tool is designed for use when prompted for a password or token. Simply background the current task or use a separate terminal tab or split to quickly get the token into your clipboard, ready for pasting.

## Usage

Call `passme` to get a TUI containing your tokens
```
> passme

Saved keys:
╭──────────────────────┬──────────────────────────────────────────────╮
│        ALIAS         │                    TOKEN                     │
├──────────────────────┼──────────────────────────────────────────────┤
│   > github_nmiguel   │   ************************************dasf   │
╰──────────────────────┴──────────────────────────────────────────────╯
```
Possible commands from the TUI
```
 Flag          Description

 a             Add a new key
 e             Edit the token under the cursor
 d             Delete the token under the cursor
 q             Return to previous screen
 enter         Copy the token under the cursor
```
Add a new token from the TUI
```
Adding new key

Alias:
> token_name

Token:
> **********
```
You can also directly copy a token's value using the copy (or c) flag (supports fuzzy finding).
```
> passme copy github
Successfully copied token with name github_nmiguel
```

## Installation

To install, use Go's `install` command:

`go install github.com/nmiguel/passme@latest`

### Dependencies

- Go version 1.20 or higher
- `xclip` or `xsel` for clipboard support

## License

Passme is licensed under the GNU GPL 3.0 License. See the `LICENSE` file for details.
