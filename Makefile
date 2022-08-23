
build:
	@go build -o build/yaml_to_json
	@GOOS=windows GOARCH=amd64 go build -o build/yaml_to_json.exe