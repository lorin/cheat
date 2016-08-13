.PHONY: run

run: cheat
	./cheat foo

cheat: main.go
	golint
	go tool vet .
	errcheck -ignore 'Close'
	go build

clean:
	rm -r foo-cheat-sheet
	rm cheat

