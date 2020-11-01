#!/bin/bash

set -e

exercise_name="$1"

mkdir "./cmd/$exercise_name"

cat > "./cmd/$exercise_name/main.go" <<EOF
package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
}
EOF
