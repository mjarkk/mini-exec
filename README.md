# `mini-exec` - A small CD (continuous deployment)
A small CD for limited environments where there are no open ports  
Mostly made to run in production docker containers as CMD command  

## Features:
1. **Run without webserver** - Can run in servers that have strict rules and can't have exposed ports
2. **Build script** - a nice build script where the steps to run the project are in defined
3. **Docker support** - Runs perfectly as underlayer for a server in a docker container
4. **1 executable and 1 config** - That's all to get this up and running

## Why?
I got annoyed of the limitations from the servers my programs need to run on and wanted a not complicated solution where i did not need 4 servers/vms/containers in total.

## Install

#### Latests release
```
wget https://github.com/mjarkk/mini-exec/releases/download/v0.1.1/release.zip
unzip release.zip
mv mini-exec /usr/bin/
```

#### From source
```
go get github.com/mjarkk/mini-exec
```

## Use

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

## TODOS/BUGS
- *todo* the documentation is minimal, a view examples would be great 
- *bug* no support for `&&`, `||`, `>` and `>>` in the shell
- [*todo* website where it's possible to directly download the executable](https://github.com/mjarkk/mini-exec/pull/4)
- *todo* no way to manually execute actions via the cli, i'm for now not sure what to do with this.
- *todo* windows support *(It probebly works on windows but there will be a view bugs like the program will see `cd C:\` as `cd C:\dir\to\project\C:\`)*
