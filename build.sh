#!/bin/bash

GOOS=darwin go build -o build/darwin/osupdate
GOOS=windows go build -o build/windows/osupdate.exe