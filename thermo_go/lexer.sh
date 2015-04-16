#!/bin/bash -e

ragel -Z lexer_go.rl 
go generate parser.go
go build -o thermo .
./thermo
ragel -V lexer_go.rl > graph.dot
if [ $(which dot) -a "$1" != "" ]; then
  dot -Tpng graph.dot > graph.png
  open graph.png
fi
