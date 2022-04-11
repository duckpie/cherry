module github.com/duckpie/cherry

go 1.18

require google.golang.org/grpc v1.45.0

require (
	github.com/golang/protobuf v1.5.2 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.26.0 // indirect
)

replace google.golang.org/genproto => ./libs/genproto
