#!/bin/bash

exercise_name=$1

if [[ $exercise_name ]];
then
  go test "./cmd/$exercise_name/..."
else
  go test ./...
fi
