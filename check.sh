#!/bin/bash

set -e

ACTION=$1

function format() {
    echo "Running gofmt ..."
    if [[ $1 == "-w" ]]; then
        gofmt -w $(find . -type f -name '*.go')
    elif [[ $1 == "-l" ]]; then
        gofmt -l $(find . -type f -name '*.go')
    elif [[ $1 == "-d" ]]; then
        gofmt -d $(find . -type f -name '*.go')
    else
        UNFORMATTED=$(gofmt -l $(find . -type f -name '*.go'))
        if [[ ! -z "$UNFORMATTED" ]]; then
            echo "The following files are not properly formatted:"
            echo "$UNFORMATTED"
            exit 1
        fi
    fi
}

function lint() {
    echo "Running golint ..."
    go install golang.org/x/lint/golint
    golint -set_exit_status ./...
}

# There is a bug if we define the same json tag in the struct.
# Disable structtag temporarily.
# see: https://github.com/golang/go/issues/40102
function vet() {
    echo  "Running go vet ..."
    (
        cd v2
        go vet -structtag=false ./...
    )
}

function unittest() {
    echo "Running go test ..."
    (
        cd v2
        go test -v -race -coverprofile=coverage.txt -covermode=atomic ./...
    )
}

function integration() {
    echo "Running integration test ..."
    cd v2
    go test -v -tags=integration -race -coverprofile=coverage.txt -covermode=atomic ./...
}

if [[ -z $ACTION ]]; then
    format
    # lint
    vet
    unittest
else
    shift
    $ACTION "$@"
fi
