#!/bin/sh
go env -w GO111MODULE=off
go install --gcflags="-N -l" github.com/alexgordon11/hello
go install github.com/alexgordon11/debugger