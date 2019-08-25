#!/bin/sh
name=${PWD##*/}
ENV=LOCAL
EXPOSED_PORT=8000

# Compile go program without debug symbols
echo "[+] Go build"
GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags "-w -s"

# Create docker image
echo "[+] Docker rm"
docker rm -f "$name"
echo "[+] Docker rmi"
docker rmi "$name"
echo "[+] Docker build"
docker build . -t "$name" --build-arg ENV=$ENV
echo "[+] Docker run"
docker run -e "ENV=$ENV" -e "VAULT_TOKEN=$VAULT_TOKEN" --name "$name" -p $EXPOSED_PORT:$EXPOSED_PORT  "$name" 
