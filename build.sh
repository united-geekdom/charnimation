#!/bin/bash
echo "Building project"
if [ -d "build" ] 
then
    rm -rf build/*
else
	mkdir build/
fi
go build -ldflags "-s -w"
echo "Copying files"
mv skimmer build/
cp -r data/* build/
cd build 
./skimmer -f $1
