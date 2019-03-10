#!/bin/bash

# Bulid the go project
GOOS=linux go build

# zip the files
zip release.zip mini-exec

echo "----------------------------"
echo "Created ./release.zip"
echo "----------------------------"
