#!/bin/bash
# Build script for vai

set -e

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Variables
BINARY_NAME="vai"
BUILD_DIR="build"
VERSION=${VERSION:-$(git describe --tags --always --dirty 2>/dev/null || echo "dev")}
LDFLAGS="-ldflags -X main.Version=${VERSION}"

echo -e "${GREEN}Building ${BINARY_NAME}...${NC}"
echo "Version: ${VERSION}"

# Create build directory
mkdir -p "${BUILD_DIR}"

# Build
go build -v ${LDFLAGS} -o "${BUILD_DIR}/${BINARY_NAME}" ./cmd/vai

if [ $? -eq 0 ]; then
    echo -e "${GREEN}Build successful!${NC}"
    echo "Binary: ${BUILD_DIR}/${BINARY_NAME}"
else
    echo -e "${RED}Build failed!${NC}"
    exit 1
fi
