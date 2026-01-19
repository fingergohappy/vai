#!/bin/bash
# Install script for vai

set -e

# Colors
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Variables
BINARY_NAME="vai"
INSTALL_DIR="${HOME}/.local/bin"

echo -e "${GREEN}Installing ${BINARY_NAME}...${NC}"

# Create install directory if it doesn't exist
mkdir -p "${INSTALL_DIR}"

# Build and install
go build -o "${INSTALL_DIR}/${BINARY_NAME}" ./cmd/vai

# Check if INSTALL_DIR is in PATH
if [[ ":$PATH:" != *":${INSTALL_DIR}:"* ]]; then
    echo -e "${YELLOW}Warning: ${INSTALL_DIR} is not in your PATH${NC}"
    echo "Add the following to your shell profile (~/.bashrc, ~/.zshrc, etc.):"
    echo "  export PATH=\"${INSTALL_DIR}:\$PATH\""
else
    echo -e "${GREEN}Installation successful!${NC}"
    echo "Run '${BINARY_NAME}' from anywhere in your terminal."
fi
