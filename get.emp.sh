#!/usr/bin/env bash
set -e

CLI_NAME="emp"
VERSION="v0.1.3"  # you could also fetch latest from GitHub API
REPO="hclpandv/emp-cli"

OS=$(uname | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)

# Map architecture
case $ARCH in
    x86_64) ARCH="amd64" ;;
    arm64|aarch64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

# download URL
URL="https://raw.githubusercontent.com/hclpandv/vikiazscan/refs/heads/main/emp-linux-amd64"

# Download binary
TMP_FILE=$(mktemp)
echo "Downloading $URL ..."
curl -L -o "$TMP_FILE" "$URL"

# Make executable
chmod +x "$TMP_FILE"

# Install to $HOME/.local/bin
INSTALL_DIR="$HOME/.local/bin"
if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
    echo "WARNING: ~/.local/bin is not in PATH"
    echo "Add it using:"
    echo '  echo "export PATH=\$HOME/.local/bin:\$PATH" >> ~/.bashrc'
fi

mv "$TMP_FILE" "$INSTALL_DIR/$CLI_NAME"

echo "$CLI_NAME installed to $INSTALL_DIR/$CLI_NAME"
echo "Run '$CLI_NAME --help' to get started."
