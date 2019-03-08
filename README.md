# `mini-exec` - A small CD (continuous deployment)
A small CD for limited environments where there are no open ports

## Goals of this project:
1. Beable to run in servers that can't host a webserver (so no fronend and only pulling data)
2. Have a build script
3. Must perfectly run as underlayer for the main app in a docker container
4. 1 executable and 1 config file for every docker container

## Why?
I got annoyed of the limitations from the servers my programs need to run on and wanted a not complicated solusion where i did not need 4 servers/vms/containers in total.

## Some other detials:
- Written in GoLang
