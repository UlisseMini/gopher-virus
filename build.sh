#!/bin/zsh

# Just in case the global aliases are not loaded
alias wgo="env GOOS=windows go build -ldflags -H=windowsgui"

# format everything first
gofmt -w -l -s .

# build payload
printf "go build payload.go\n"
cd ./payload
wgo payload.go
mv payload.exe .. 2>/dev/null
cd ..

# build deploy
printf "go build deploy.go\n"
cd ./deploy
wgo deploy.go
mv deploy.exe .. 2>/dev/null
cd ..
