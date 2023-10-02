#!/bin/bash
function show_progress() {
    local duration=$1
    local bar_length=10
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
# sleep 2
cd ./cmd/master-sec && go build -o ../../bin/porty.exe
# show_progress 1
# echo "" 
# echo "Testing App..."
# cd ../../pkg
# go test -v ./...
# cd ..
# show_progress 1
# echo ""
# echo "completed!"
./bin/porty.exe

