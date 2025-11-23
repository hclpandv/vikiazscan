#!/usr/bin/env bash
set -e

CLI_NAME="emp"
REPO="hclpandv/vikiazscan"

# Always install into user-writable location
INSTALL_DIR="$HOME/.local/bin"
mkdir -p "$INSTALL_DIR"

# Arch/OS detection
ARCH=$(uname -m)
case $ARCH in
    x86_64) ARCH="amd64" ;;
    arm64|aarch64) ARCH="arm64" ;;
    *) echo "Unsupported architecture: $ARCH"; exit 1 ;;
esac

OS=$(uname | tr '[:upper:]' '[:lower:]')

# Download URL for Linux binary
URL="https://raw.githubusercontent.com/hclpandv/vikiazscan/refs/heads/main/emp-linux-amd64"

# Download to temp file
TMP_FILE=$(mktemp)
echo "Downloading $URL ..."
curl -fsSL -o "$TMP_FILE" "$URL"

# Make executable
chmod +x "$TMP_FILE"

# Move WITHOUT SUDO
mv "$TMP_FILE" "$INSTALL_DIR/$CLI_NAME"

echo "EMP installed to: $INSTALL_DIR/$CLI_NAME"

# PATH check
if ! echo "$PATH" | grep -q "$INSTALL_DIR"; then
    echo
    echo "⚠️  ~/.local/bin is NOT in your PATH."
    echo "Run this to fix:"
    echo 'echo "export PATH=$HOME/.local/bin:$PATH" >> ~/.bashrc'
    echo "Then restart your shell."
    echo
fi

echo "Run: emp --version"
