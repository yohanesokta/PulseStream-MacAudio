#!/bin/bash
echo  "Checking Docker..."
if command -v docker >/dev/null 2>&1; then
    echo "."
    docker --version
else
    echo "Docker is NOT installed. Please Install Docker"
    exit 1
fi

