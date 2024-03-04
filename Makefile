generate:
	GO_POST_PROCESS_FILE="gofmt -w" openapi-generator \
		generate --enable-post-process-file -i openapispec.json -g go --package-name tabular  \
		--git-host github.com --git-user-id tabular-io --git-repo-id tabular-sdk-go \
		--additional-properties=isGoSubmodule=true,withGoMod=false \
		-o tabular/

replace_version: generate
	sed -i'' -e 's/var Version.*/var Version = \"${VERSION_TAG}\"/g' tabular/version.go


build: replace_version