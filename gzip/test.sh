#!/bin/bash -x

go build .

# pre-heat things
./main > /dev/null

# time it
time ./main

# pre-heat again
gzip -9kf test.txt > /dev/null

# time the gzip
time gzip -9kf test.txt

