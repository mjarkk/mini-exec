#!/bin/bash

cd js
yarn build
cd ..
go generate
go install
mini-exec

