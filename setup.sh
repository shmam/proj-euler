#!/bin/bash

if [ -z "$1" ]; then
    echo "Please provide a number for the problem."
    exit 1
fi

number=$1
echo "🫵 You provided the number: $number"

mkdir -p "$number"
echo "📂 Directory $number created."

curl -o "$number/problem.md" "https://projecteuler.net/minimal=$number"
echo "📝 Problem statement saved to $number/problem.md."

cat <<EOL > "$number/main.go"
package main

import "os"
import "fmt"

func main() {
    fmt.Println("solution goes here!")
    os.Exit(1)
}
EOL

echo "⚙️ Go file created at $number/main.go."

cd "$number"
go mod init "$number"
echo "👍 Go module initialized in $number."