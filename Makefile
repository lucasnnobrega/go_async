all:
	go build -o ./build ./sync/*
	go build -o ./build ./chan/*
	go build -o ./build ./src/*

build:
	go build -o ./build ./sync/*
	go build -o ./build ./chan/*
	go build -o ./build ./src/*

run:
	go run ./sync/*
	go run ./chan/*
	go run ./src/*

clean:
	rm ./build/json*
