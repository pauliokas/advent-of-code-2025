.PHONY: test
test: days/*
	go test -v $(addprefix ./, $^)

.PHONY: bench
bench: days/*
	go test -bench=. -v $(addprefix ./, $^)
