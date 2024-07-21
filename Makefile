.ONESHELL:
.SHELLFLAGS += -e
.SHELLFLAGS += -o pipefail

gorelease:
	cd go
	goreleaser build --snapshot --clean
	rm -f dist/artifacts.json dist/config.yaml dist/metadata.json