#!/bin/bash

RUN_NAME="go_schedule"

mkdir output output/${RUN_NAME}_log
cp -r config output/
export GO111MODULE=on
# dev
# go build -a -o output/${RUN_NAME} -ldflags "-w"

# prod
go build -a -o output/${RUN_NAME}