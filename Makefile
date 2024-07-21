# .ONESHELL:
# .SHELLFLAGS += -e

gorelease:
	cd go && goreleaser build --snapshot --clean && rm -f dist/artifacts.json dist/config.yaml dist/metadata.json

centralize_release_files:
	rm -rf dist
	mkdir -p dist
	@mv -f go/dist/* dist/ 2>/dev/null || echo "> go: no files to move"