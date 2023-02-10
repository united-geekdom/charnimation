#!/bin/bash
echo "Building project"
if [ -d "data" ] 
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
./skimmer -f $1 -w $2 -h $3
