# `mini-exec` - A small CD (continuous deployment)
A small CD for limited environments where there are no open ports  
Mostly made to run in production docker containers as CMD command  

## Goals of this project:
1. Beable to run in servers that can't host a webserver (so no fronend and only pulling data)
2. Have a build script
3. Must perfectly run as underlayer for the main app in a docker container
4. 1 executable and 1 config file to get this up and working

## Why?
I got annoyed of the limitations from the servers my programs need to run on and wanted a not complicated solusion where i did not need 4 servers/vms/containers in total.

## Install
```
wget https://github.com/mjarkk/haproxy-check-api/releases/download/0.1/release.zip
unzip release.zip
mv mini-exec /usr/bin/
```

## Use
*NOTE: this program is not tested yet in a production envourment*  

#### 1. Make a build script `.miniex`
To add a automated build process we need a `.miniex` file.  
The contents of a `.miniex` file will look simaliar to a bash file,  
although it will work more like a dockerfile.  
Under here is a examle `.miniex` file.  
```bash
# A example miniex file for a go project with a javascript frontend

# Build go to binary
go build -o ./my-webserver

# Build the frondend javascript files
cd frontend
npm ci
npm run build
cd ..

# Run the binary
# The FINAL command is like the docker CMD it suposed to be a command that runs forever.
FINAL ./my-webserver
```

#### 2. Run it
Run: `$ mini-exec`  
This will run all steps in `.miniex` and if it no steps failed it will run the command after FINAL

#### 3. Updates
mini-exec will do a `git pull` every 2 minutes and if there are updates it will re-run the content of the `.miniex` file.  
When it reaches the `FINAL ...` command without any errors it stops what was running from the last miniex file and start the new command.
