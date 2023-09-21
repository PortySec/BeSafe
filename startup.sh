#!/bin/bash
function show_progress() {
    local duration=$1
    local bar_length=20
    local sleep_delay=$((duration / bar_length))

    for ((i = 0; i <= bar_length; i++)); do
        printf "[%-${bar_length}s] %d%%" "$(printf '#%.0s' $(seq 1 $i))" "$((i * 100 / bar_length))"
        sleep "$sleep_delay"
        printf "\r"
    done
}
clear

echo "Compiling tools..."

# Simulating some task being performed
sleep 2
cd ./cmd/passguard && go build -o ../../bin/passguard.exe
cd -
cd ./cmd/scanner && go build -o ../../bin/scanner.exe
show_progress 2
echo "" 
echo "Testing App..."
cd ../..
go test -v ./...
show_progress 2
echo ""
go build -o porty.exe
echo "completed!"

