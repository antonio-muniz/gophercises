#!/bin/bash

exercise_name="$1"

go build -i -o "./bin/$exercise_name" "./cmd/$exercise_name"

shift
"./bin/$exercise_name" "$@"
