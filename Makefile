PREFIX=/usr/local
VERSION=v$(shell awk '/app.Version/ { print $$3 }' main.go)

version:
	git tag -f $(VERSION)

format:
	gofmt -s -w .

build: format
	go build -trimpath -buildmode=pie -mod=readonly -modcacherw -ldflags="-s -w"

install: build
	mkdir -p $(PREFIX)/bin
	mkdir -p /usr/share/servus
	cp -f icons/* /usr/share/servus
	cp -f i3-utils $(PREFIX)/bin

uninstall:
	rm -f $(PREFIX)/bin
