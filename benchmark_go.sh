#!/bin/bash -eu

count=100000
for ((i=0; i < $count; i++)); do
    ./gomd5 "$i"
done
