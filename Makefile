clean:
	rm *.test

run:
	go run .

.PHONY: clean run
.SILENT: clean run