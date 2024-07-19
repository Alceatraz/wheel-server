@echo off

go build -o bin\server-x64.exe

set GOOS=linux

go build -o bin\server
