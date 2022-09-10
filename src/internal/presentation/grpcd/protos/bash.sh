#!/bin/bash

protoc --proto_path=. --go_out=../pbs --go-grpc_out=require_unimplemented_servers=false:../pbs *.proto

#ls ../pbs/balance.pb.go | xargs -n1 -ix bash -c 'sed s/,omitempty// x > x.tmp && mv x{.tmp,}'
