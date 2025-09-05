module github.com/sawpanic/CProtocol

go 1.21

require (
	github.com/gorilla/websocket v1.5.3
	github.com/redis/go-redis/v9 v9.5.1
	github.com/rs/zerolog v1.31.0
	github.com/sony/gobreaker v0.5.0
	github.com/spf13/cobra v1.8.0
)

require (
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.12.0 // indirect
)

replace github.com/cryptoedge/internal => ./internal
