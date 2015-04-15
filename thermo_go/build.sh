#!/bin/bash

go generate parser.go
go build -o thermo .
./thermo
