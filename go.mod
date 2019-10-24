module github.com/zahfox/gourd

go 1.13

replace github.com/zahfox/gourd => ./pkg

require (
	github.com/coreos/go-systemd v0.0.0-20190719114852-fd7a80b32e1f
	github.com/pkg/errors v0.8.1
	github.com/urfave/cli v1.22.1
)
