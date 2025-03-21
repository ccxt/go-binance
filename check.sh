#!/bin/bash

set -e

ACTION=$1

function format() {
    echo "Running gofmt ..."
    if [[ $1 == "-w" ]]; then
        gofmt -w $(find . -type f -name '*.go' -not -path "./vendor/*")
    elif [[ $1 == "-l" ]]; then
        gofmt -l $(find . -type f -name '*.go' -not -path "./vendor/*")
    else
        test -z "$(gofmt -l $(find . -type f -name '*.go' -not -path "./vendor/*"))"
    fi
}

function lint() {
    echo "Running golint ..."
    go install golang.org/x/lint/golint
    golint -set_exit_status ./...
}

function vet() {
    echo  "Running go vet ..."
    (
        cd v2
        go vet ./...
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
