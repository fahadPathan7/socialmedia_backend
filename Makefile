protogen-user:
	protoc -I. \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=logtostderr=true,grpc_api_configuration=./proto/user/user.http.yaml:. \
	--swagger_out=logtostderr=true,grpc_api_configuration=./proto/user/user.http.yaml:. \
	./proto/user/user.proto

protogen-post:
	protoc -I. \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=logtostderr=true,grpc_api_configuration=./proto/post/post.http.yaml:. \
	--swagger_out=logtostderr=true,grpc_api_configuration=./proto/post/post.http.yaml:. \
	./proto/post/post.proto