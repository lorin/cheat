.PHONY: run

run: clean cheat
	./cheat foo

cheat: main.go
	golint
	go tool vet .
	errcheck -ignore 'Close'
	go build

clean:
	rm -rf foo-cheat-sheet
	rm cheat

