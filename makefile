run:
	./bin/heartack
install:
	go get .
	go install
full:
	make install
	make run