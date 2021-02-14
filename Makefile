PREFIX=/usr/local

build:
	gofmt -w .
	go build

install: build
	mkdir -p $(PREFIX)/bin
	mkdir -p /usr/share/servus
	cp -f icons/* /usr/share/servus
	cp -f i3-utils $(PREFIX)/bin

uninstall:
	rm -f $(PREFIX)/bin
