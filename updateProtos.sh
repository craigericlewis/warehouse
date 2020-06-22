#!/bin/bash

for f in $(find ./server/protos -name '*.proto');
do
	echo "Processing $f"
    protoc --go_out=plugins=grpc:. $f
done
