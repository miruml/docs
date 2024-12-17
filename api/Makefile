all:
	openapi-generator-cli generate -i main.yaml -g go-server -o ./all/server
	openapi-generator-cli generate -i main.yaml -g go -o ./all/client
	cp ./all/server/api/openapi.yaml ./openapi.yaml

models: 
	openapi-generator-cli generate -i main.yaml -g go-server --global-property models -o ./models/server
	openapi-generator-cli generate -i main.yaml -g go --global-property models -o ./models/client

clean:
	rm -rf ./all ./models

rebuild: clean all
