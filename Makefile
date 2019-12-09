build:
	go build

run: clean build
	./vocability

clean:
	rm vocability
