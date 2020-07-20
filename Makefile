all:
	go build -o ./build ./compress/*
	go build -o ./build ./sync/*
	go build -o ./build ./chan/*
	go build -o ./build ./src/*

build:
	go build -o ./build ./compress/*
	go build -o ./build ./sync/*
	go build -o ./build ./chan/*
	go build -o ./build ./src/*

run:
	go run ./compress/*
	go run ./sync/*
	go run ./chan/*
	go run ./src/*

clean:
	rm ./build/json*
