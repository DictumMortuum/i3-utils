install: build
	cp i3-utils ~/.local/bin/

build:
	go build

uninstall:
	rm ~/.local/bin/i3-utils
