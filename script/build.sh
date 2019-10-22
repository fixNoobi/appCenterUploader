#!/usr/bin/env sh

go build
GOARCH=amd64
for GOOS in darwin linux; do
     export GOOS $GOARCH
     go build -v -o appCenterUploader-$GOOS-$GOARCH ./cmd
done