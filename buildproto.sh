protoc -I proto/ proto/payment.proto --go_out=plugins=grpc:proto
ls proto/*.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'