PREFIX=/usr/local

build:
	go build

install: build
	mkdir -p $(PREFIX)/bin
	cp -f i3-utils $(PREFIX)/bin

uninstall:
	rm -f $(PREFIX)/bin
