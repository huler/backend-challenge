#!/bin/bash

rm -rf dist
mkdir dist

(cd ./postresults/ && ./build.sh)

(cd ./getresults/ && ./build.sh)

