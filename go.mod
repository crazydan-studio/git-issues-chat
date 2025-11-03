module git-issues-chat

go 1.24.0

toolchain go1.24.9

require (
	github.com/mattn/go-sqlite3 v1.14.32
	github.com/webui-dev/go-webui/v2 v2.4.2
	golang.org/x/oauth2 v0.32.0
)

replace github.com/webui-dev/go-webui/v2 => ./deps/go-webui/v2
