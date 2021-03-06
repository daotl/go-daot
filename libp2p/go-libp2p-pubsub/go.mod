module github.com/daotl/go-libp2p-pubsub

go 1.13

replace github.com/libp2p/go-libp2p-pubsub => ./

require (
	github.com/benbjohnson/clock v1.1.0
	github.com/gogo/protobuf v1.3.2
	github.com/ipfs/go-log v1.0.5
	github.com/klauspost/compress v1.12.3 // indirect
	github.com/libp2p/go-libp2p-blankhost v0.2.0
	github.com/libp2p/go-libp2p-connmgr v0.2.4
	github.com/libp2p/go-libp2p-core v0.11.0
	github.com/libp2p/go-libp2p-discovery v0.6.0
	github.com/libp2p/go-libp2p-pubsub v0.0.0-00010101000000-000000000000
	github.com/libp2p/go-libp2p-swarm v0.8.0
	github.com/libp2p/go-msgio v0.1.0
	github.com/multiformats/go-multiaddr v0.4.1
	github.com/whyrusleeping/timecache v0.0.0-20160911033111-cfcb2f1abfee
	go.uber.org/zap v1.19.1 // indirect
)
