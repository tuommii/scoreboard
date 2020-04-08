
all: run

run: build
	./run

build:
	go build -o run cmd/sgoreboard/main.go

# Runs test and bench: both go test -bench . ./...
# Now only tests
test:
	go test -v ./...

# Runs without other tests
bench:
	go test -run=XXX -bench . ./...
