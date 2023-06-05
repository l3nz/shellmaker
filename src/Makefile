all:
	rm -fr shm ; go build -race -o shm

test:
	go build -o shm && \
	#go test -v ./cfgfile  && \
	go test -v . && \
	echo "ALL TESTS WENT WELL" 

