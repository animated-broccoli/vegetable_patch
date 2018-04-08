CC=go


all: get build run

get: 
	$(CC) get

build:
	$(CC) build

run:
	./vegetable_patch


