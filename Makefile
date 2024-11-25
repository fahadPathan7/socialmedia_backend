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

protogen-comment:
	protoc -I. \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=logtostderr=true,grpc_api_configuration=./proto/comment/comment.http.yaml:. \
	--swagger_out=logtostderr=true,grpc_api_configuration=./proto/comment/comment.http.yaml:. \
	./proto/comment/comment.proto

protogen-react:
	protoc -I. \
	--go_out=. \
	--go-grpc_out=. \
	--grpc-gateway_out=logtostderr=true,grpc_api_configuration=./proto/react/react.http.yaml:. \
	--swagger_out=logtostderr=true,grpc_api_configuration=./proto/react/react.http.yaml:. \
	./proto/react/react.proto