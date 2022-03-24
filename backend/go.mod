module github.com/colbylwilliams/clutch-teamcloud-gateway/backend

go 1.16

require (
	github.com/bufbuild/buf v0.56.0
	github.com/colbylwilliams/teamcloud-go/teamcloud v0.0.0-20220324192048-86b286d44b83
	github.com/envoyproxy/protoc-gen-validate v0.6.7
	github.com/fullstorydev/grpcurl v1.8.6
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.9.0
	github.com/lyft/clutch/backend v0.0.0-20220323165035-8f7d2f4d09df
	github.com/lyft/clutch/tools v0.0.0-20220323165035-8f7d2f4d09df
	github.com/shurcooL/vfsgen v0.0.0-20200824052919-0d455de96546
	github.com/stretchr/testify v1.7.1
	github.com/uber-go/tally/v4 v4.1.1
	go.uber.org/zap v1.21.0
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106
	google.golang.org/grpc v1.45.0
	google.golang.org/grpc/cmd/protoc-gen-go-grpc v1.2.0
	google.golang.org/protobuf v1.27.1
)
