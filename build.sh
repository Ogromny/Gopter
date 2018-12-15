#!/bin/bash
go build -o builder.elf builder/main.go
go build -o stub.elf stub/main.go