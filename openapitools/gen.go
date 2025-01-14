package openapitools

//go:generate java -jar openapi-generator-cli-7.9.0.jar generate -i ../openapi3.yaml -g go -o gen/client --enable-post-process-file -c .client.yaml
//go:generate java -jar openapi-generator-cli-7.9.0.jar generate -i ../openapi3.yaml -g go-server -o gen/server --enable-post-process-file -c .server.yaml
