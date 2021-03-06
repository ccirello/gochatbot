GCBFLAGS=all

all: vendor
	GO15VENDOREXPERIMENT=1 go build -tags "$(GCBFLAGS)"

docker: vendor
	GOOS=linux GO15VENDOREXPERIMENT=1 go build -tags "all" -o gochatbot-container
	docker build -t ccirello/gochatbot .
	rm -f gochatbot-container

vendor:
	go get github.com/Masterminds/glide
	GO15VENDOREXPERIMENT=1 $(GOPATH)/bin/glide -q install

clean:
	rm -rf vendor/ gochatbot

novendor:
	cat glide.yaml | grep -i '\- package' | awk '{ print $$3 }' | xargs go get -u || exit 0

quickstart: novendor
	go build -tags "$(GCBFLAGS)"

release: vendor
	rm -f gochatbot-linux gochatbot-darwin gochatbot-windows gochatbot-linux.gz gochatbot-darwin.gz gochatbot-windows.gz
	GOOS=linux GO15VENDOREXPERIMENT=1 go build -tags "all" -o gochatbot-linux
	gzip -f gochatbot-linux
	GOOS=darwin GO15VENDOREXPERIMENT=1 go build -tags "all" -o gochatbot-darwin
	gzip -f gochatbot-darwin
	GOOS=windows GO15VENDOREXPERIMENT=1 go build -tags "all" -o gochatbot-windows
	gzip -f gochatbot-windows
