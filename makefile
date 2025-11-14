run:
	go test ./... -covermode=count -coverprofile="coverage.out" fmt
	go tool cover -html="coverage.out"
	go tool cover -func="coverage.out" -o="coverage.out"

badge: run
	gobadge -filename="coverage.out" -yellow=30 -green=70 -target="README.md"