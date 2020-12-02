schema:
	@read -p "Enter Schema name: " name; \
		ent init $$name

generate:
	go generate ./ent

mac:
	env GOOS=darwin GOARCH=amd64 go build

linux:
	env GOOS=linux GOARCH=amd64 go build

windows:
	env GOOS=windows GOARCH=amd64 go build