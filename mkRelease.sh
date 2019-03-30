#!/bin/bash

# Build the web related files
cd js
yarn
yarn build
cd ..
go generate

# Bulid the go project
GOOS=linux go build

# zip the files
zip release.zip mini-exec

echo "----------------------------"
echo "Created ./release.zip"
echo "----------------------------"
