.PHONY: builder stub hello clean

builder:
	go build -o builder ./cmd/builder

stub:
	go build -o stub ./cmd/stub

hello:
	$(CC) -o hello hello.c

clean:
	rm -f builder stub hello binded
