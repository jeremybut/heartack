run:
	$(info ************  Compile and run Go program ************)
	./bin/heartack

get:
	$(info ************  Download and install packages and dependencies ************)
	go get ./src/heartack

install:
	$(info ************  Compile and install packages and dependencies ************)
	go install ./src/heartack

bower:
	$(info ************  Bower installs packages to bower_components ************)
	bower install

clean:
	$(info ************  Remove binaries, packages, components files. ************)
	rm -rf ./bin
	rm -rf ./pkg
	rm -rf ./src/github.com
	rm -rf ./src/gopkg.in
	rm -rf ./static/vendor/components

full:
	make get
	make bower
	make install
	make run