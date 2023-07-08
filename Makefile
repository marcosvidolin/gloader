build:
	go build . 
	
run: build
	./gloader -s /path/to/scenarios.yml