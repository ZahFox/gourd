module github.com/zahfox/gourd

go 1.13

replace github.com/zahfox/gourd => ./pkg

require (
	github.com/c-bata/go-prompt v0.2.3
	github.com/google/uuid v1.1.1
	github.com/mattn/go-tty v0.0.3 // indirect
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.8.1
	github.com/pkg/term v0.0.0-20190109203006-aa71e9d9e942 // indirect
	github.com/sirupsen/logrus v1.4.2
	github.com/spf13/viper v1.5.0
	github.com/ttacon/chalk v0.0.0-20160626202418-22c06c80ed31
	github.com/ugorji/go/codec v1.1.7
	github.com/urfave/cli v1.22.1
	gopkg.in/src-d/go-git.v4 v4.13.1
)
