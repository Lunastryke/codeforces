#!/bin/bash

# Usage: ./run_test.sh problems/contest_1234/problem_a/main.go

SOLUTION=$1
INPUT=${SOLUTION%/*}/input.txt
OUTPUT=${SOLUTION%/*}/output.txt

go run $SOLUTION < $INPUT > result.txt

if diff -q result.txt $OUTPUT > /dev/null; then
    echo "Test passed!"
else
    echo "Test failed!"
    diff result.txt $OUTPUT
fi

rm result.txt
