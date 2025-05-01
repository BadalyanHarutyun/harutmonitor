#!/bin/bash

set -e

BINARY_NAME="harutmonitor"
TMP_PATH="./$BINARY_NAME"
INSTALL_PATH="/usr/local/bin/$BINARY_NAME"

echo "🔐 Setting executable permissions..."
chmod +x "$TMP_PATH"

echo "🚚 Moving $BINARY_NAME to $INSTALL_PATH..."
sudo mv "$TMP_PATH" "$INSTALL_PATH"

echo "✅ Installed $BINARY_NAME to $INSTALL_PATH"
