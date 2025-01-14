openapi:
	openapi-overlay apply overlay.yaml ../original.yaml > openapi.yaml

gen-client:
	GO_POST_PROCESS_FILE='gofmt -w' java -jar openapi-generator-cli-7.9.0.jar generate -i openapi.yaml -g go -o gen/client --enable-post-process-file -c .client.yaml

gen-server:
	GO_POST_PROCESS_FILE='gofmt -w' java -jar openapi-generator-cli-7.9.0.jar generate -i openapi.yaml -g go-server -o gen/server --enable-post-process-file -c .server.yaml
