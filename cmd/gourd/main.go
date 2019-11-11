// Gourd is my personal linux configuration tool.
package main

import (
	"fmt"
	"net"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/urfave/cli"

	"github.com/zahfox/gourd/internal/gourd/config"
	"github.com/zahfox/gourd/pkg/client"
	"github.com/zahfox/gourd/pkg/distro"
	"github.com/zahfox/gourd/pkg/misc"
	"github.com/zahfox/gourd/pkg/utils"
)

func main() {
	app := cli.NewApp()
	configureAppInfo(app)
	configureAppCommands(app)
	configureAppAction(app)
	err := app.Run(os.Args)
	checkerr(err)
}

func configureAppInfo(app *cli.App) {
	app.Name = "gourd"
	app.Usage = "linux configuration tool"
}

func configureAppCommands(app *cli.App) {
	ipFlags := []cli.Flag{
		cli.StringFlag{
			Name: "host",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:  "ip",
			Usage: "Looks up the IP address for a particular host",
			Flags: ipFlags,
			Action: func(c *cli.Context) error {
				var host string
				arg := c.Args().First()
				if len(arg) > 0 {
					host = arg
				} else {
					host = c.String("host")
				}

				if len(host) < 1 {
					host = "localhost"
				}

				ip, err := net.LookupIP(host)
				if err != nil {
					return errors.Wrap(err, "Hostname error")
				}

				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}

				return nil
			},
		},
		{
			Name:  "client",
			Usage: "Starts a new gourd client",
			Action: func(c *cli.Context) error {
				client := client.NewClient(config.GetSocketPath())
				client.Run()

				go func() {
					for i := 10; i > 0; i-- {
						client.Ping()
						client.Ping()
						client.Ping()
						client.Echo("ECHO")
					}
					client.Stop()
				}()

				utils.LogInfo("Waiting for gourd client to exit")
				client.Wait()
				client.Exit()
				return nil
			},
		},
		{
			Name:  "install",
			Usage: "Installs a package using the distribution package manager",
			Action: func(c *cli.Context) error {
				distro.GetDistro().Install(c.Args()...)
				return nil
			},
		},
		{
			Name:  "uninstall",
			Usage: "Uninstalls a package using the distribution package manager",
			Action: func(c *cli.Context) error {
				distro.GetDistro().Uninstall(c.Args()...)
				return nil
			},
		},
		{
			Name:  "misc",
			Usage: "Used to test miscellaneous experiments",
			Action: func(c *cli.Context) error {
				misc.Run()
				return nil
			},
		},
	}
}

func configureAppAction(app *cli.App) {
	app.Action = func(c *cli.Context) error {
		path, err := exec.LookPath("ls")
		checkerr(err)

		success, err := utils.UserCanExec(path)
		checkerr(err)

		if success {
			utils.LogInfof("$USER can exec %s", path)
		}

		return nil
	}
}

func checkerr(err error) {
	if err != nil {
		utils.LogFatal(err)
	}
}
