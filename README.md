## Go - build
The Go example requires a plugin to the protocol buffer compiler, so it is not
build with all the other examples.  See:

    https://github.com/golang/protobuf

for more information about Go protocol buffer support.

First, install the Protocol Buffers compiler (protoc).

Then, install the Go Protocol Buffers plugin ($GOPATH/bin must be in your $PATH
for protoc to find it):

    go get github.com/golang/protobuf/protoc-gen-go

Build the Go samples in this directory with "make go".  This creates the
following executable files in the current directory:

    server/server_go    client/client_go

and, one more go file named 'calc.pb.go' in proto directory.
