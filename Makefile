.PHONY: test
test:
	go test ./days/...

.PHONY: bench
bench:
	go test -bench=. -benchtime=2s -cpu=1 ./days/...
