build-win:
	go build -o ./bin/app.exe ./cmd/app/

build-linux:
	go build -o ./bin/app ./cmd/app

run:
	make build-win
	./bin/app.exe

