#!/bin/bash

if [ "$#" -ne 1 ]; then
    echo "Usage: $0 <filename>"
    exit 1
fi

grep -oP "mul\(\d+,\d+\)" $1 | while read -r line; do
    values=$(echo "$line" | grep -oP "\d+")
    readarray -t numbers <<< "$values"
    factor=1
    for num in "${numbers[@]}"; do
        factor=$((factor*num))
    done
    sum=$((sum+factor))
    echo "Part1: " $sum
done
