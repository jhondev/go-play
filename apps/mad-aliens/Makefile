run:
	@go run cli/main.go run -p ./pkg/simulation/testdata/world_map_15.txt -n 10 -m 10

run-3-cities:
	@go run cli/main.go run -p ./pkg/simulation/testdata/world_map_battles.txt -n 2 -m 3

build:
	@go build -o mad ./cli

test:
	@go test ./...