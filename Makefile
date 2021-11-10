PREFIX=/usr/local
VERSION=v$(shell awk '/app.Version/ { print $$3 }' main.go)

version:
	git tag -f $(VERSION)

build:
	gofmt -s -w .
	$(info $(GOFLAGS))
	go build

install: build
	mkdir -p $(PREFIX)/bin
	mkdir -p /usr/share/servus
	cp -f icons/* /usr/share/servus
	cp -f i3-utils $(PREFIX)/bin

uninstall:
	rm -f $(PREFIX)/bin
