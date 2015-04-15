#!/bin/bash -e

ragel -Z lexer_go.rl 
go build -o hg .
./hg
ragel -V lexer_go.rl > hwg.dot
if [ $(which dot) -a "$1" != "" ]; then
  dot -Tpng hwg.dot > hwg.png
open hwg.png
fi
