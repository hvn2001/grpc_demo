1. Build out XX.pb.go: 

protoc filepath --go_out=plugins=grpc:.

2. Install grpc in Go - see GO Modules, go.mod 
(https://github.com/grpc/grpc-go):

"google.golang.org/grpc"  -->  go get -u google.golang.org/grpc

3. undefined: os.ErrDeadlineExceeded --> 

go get -u github.com/golang/protobuf/{proto,protoc-gen-go}

4. go run server/server.go: 

generate some stuffs go.mod --> go.sum ?