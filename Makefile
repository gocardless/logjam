VERSION=0.1.0

.PHONY: build rpm deb

build:
	go build -o logjam *.go

rpm deb: build
	fpm -s dir -t $@ -n logjam -v $(VERSION) \
		--description 'a log shipping tool' \
		--maintainer 'GoCardless Engineering <engineering@gocardless.com>' \
		--url 'https://github.com/gocardless/logjam' \
		logjam=/usr/bin/logjam
