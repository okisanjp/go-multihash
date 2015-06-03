#!/bin/bash -eu

count=10
for ((i=0; i < $count; i++)); do
    echo -n "$i" | md5sum
done
