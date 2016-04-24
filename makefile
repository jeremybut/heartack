run:
	$(info ************  Compile and run Go program ************)
	./bin/heartack

get:
	$(info ************  Download and install packages and dependencies ************)
	go get ./src/heartack

install:
	$(info ************  Compile and install packages and dependencies ************)
	go install ./src/heartack

full:
	make get
	make install
	make run