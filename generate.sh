#!/bin/bash

wget https://vrchatapi.github.io/specification/openapi.yaml -O openapi.yaml

go install github.com/mayocream/openapi-codegen@latest

openapi-codegen -i ./openapi.yaml -o . -p vrchat
