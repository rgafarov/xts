#!/bin/bash

mkdir -p ./build && cd build
 
cmake ../ -G Ninja

NUM_JOBS=$(( ($(nproc || grep -c ^processor /proc/cpuinfo) + 1) / 2 ))

ninja -j $NUM_JOBS
