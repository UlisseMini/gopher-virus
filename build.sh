#!/bin/zsh

# Just in case the global aliases are not loaded
alias wgo="env GOOS=windows go build -ldflags -H=windowsgui"

# format and build payload
printf "go build payload.go\n"
cd ./payload
go fmt payload.go 2>/dev/null
wgo payload.go
mv payload.exe .. 2>/dev/null
cd ..

# format and build deploy
printf "go build deploy.go\n"
cd ./deploy
go fmt deploy.go 2>/dev/null
wgo deploy.go
mv deploy.exe .. 2>/dev/null
cd ..
